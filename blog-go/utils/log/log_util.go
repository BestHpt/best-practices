/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package log

import (
	"best-practics/common"
	"best-practics/common/consts"
	"context"
	"go.uber.org/zap"
)

//type LogWrapper struct {
//	logger *zap.Logger
//}
//
//var Log LogWrapper

func Debug(tag string, fields ...zap.Field) {
	common.Logger.Debug(tag, fields...)
}

func DebugF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	common.Logger.Debug(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Info(tag string, fields ...zap.Field) {
	common.Logger.Info(tag, fields...)
}

func InfoF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	common.Logger.Info(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Warn(tag string, fields ...zap.Field) {
	common.Logger.Warn(tag, fields...)
}

func WarnF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	common.Logger.Warn(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Error(tag string, fields ...zap.Field) {
	common.Logger.Error(tag, fields...)
}

func ErrorF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	common.Logger.Error(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Fatal(tag string, fields ...zap.Field) {
	common.Logger.Fatal(tag, fields...)
}

func FatalF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	common.Logger.Fatal(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}
