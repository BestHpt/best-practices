/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package log

import (
	"best-practics/common/consts"
	trace2 "best-practics/common/trace"
	"context"
	"go.uber.org/zap"
)

type LogWrapper struct {
	ZapLogger *zap.Logger
}

var wLog *LogWrapper

func Debug(tag string, fields ...zap.Field) {
	wLog.ZapLogger.Debug(tag, fields...)
}

func DebugF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	wLog.ZapLogger.Debug(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Info(tag string, fields ...zap.Field) {
	wLog.ZapLogger.Info(tag, fields...)
}

func InfoF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	wLog.ZapLogger.Info(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Warn(tag string, fields ...zap.Field) {
	wLog.ZapLogger.Warn(tag, fields...)
}

func WarnF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	wLog.ZapLogger.Warn(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Error(tag string, fields ...zap.Field) {
	wLog.ZapLogger.Error(tag, fields...)
}

func ErrorF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	wLog.ZapLogger.Error(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Fatal(tag string, fields ...zap.Field) {
	wLog.ZapLogger.Fatal(tag, fields...)
}

func FatalF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	wLog.ZapLogger.Fatal(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}
