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

type LogWrapper struct {
	logger *zap.Logger
}

var Log LogWrapper

func Debug(tag string, fields ...zap.Field) {
	Log.logger.Debug(tag, fields...)
}

func DebugF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	Log.logger.Debug(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Info(tag string, fields ...zap.Field) {
	Log.logger.Info(tag, fields...)
}

func InfoF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	Log.logger.Info(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Warn(tag string, fields ...zap.Field) {
	Log.logger.Warn(tag, fields...)
}

func WarnF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	Log.logger.Warn(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Error(tag string, fields ...zap.Field) {
	Log.logger.Error(tag, fields...)
}

func ErrorF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	Log.logger.Error(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}

func Fatal(tag string, fields ...zap.Field) {
	Log.logger.Fatal(tag, fields...)
}

func FatalF(ctx context.Context, tag string, fields ...zap.Field) {
	trace := ctx.Value(consts.TraceKey).(*common.Trace)
	Log.logger.Fatal(tag,
		append(fields, zap.String("trace_id", trace.TraceId), zap.Int("user_id", trace.UserId))...,
	)
}
