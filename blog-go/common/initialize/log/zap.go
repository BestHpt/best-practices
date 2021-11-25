/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package log

import (
	"best-practics/common/config"
	"best-practics/utils"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap(){
	if ok, _ := utils.PathExists(config.ConfigCenter.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.ConfigCenter.Zap.Director)
		_ = os.Mkdir(config.ConfigCenter.Zap.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", config.ConfigCenter.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", config.ConfigCenter.Zap.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", config.ConfigCenter.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", config.ConfigCenter.Zap.Director), errorPriority),
	}
	//logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller(), zap.AddCallerSkip(1))

	wLog = &LogWrapper{ZapLogger: logger}
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (zapConfig zapcore.EncoderConfig) {
	zapConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.ConfigCenter.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case config.ConfigCenter.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		zapConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.ConfigCenter.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		zapConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.ConfigCenter.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		zapConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.ConfigCenter.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		zapConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return zapConfig
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.ConfigCenter.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05"))
}

//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error
func GetWriteSyncer(file string) zapcore.WriteSyncer {
	// 每小时一个文件
	logf, _ := rotatelogs.New(file +".%Y%m%d%H",
		rotatelogs.WithLinkName(file),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)

	if config.ConfigCenter.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logf))
	}
	return zapcore.AddSync(logf)
}

