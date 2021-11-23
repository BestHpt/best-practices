/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package log

import (
	"best-practics/common"
	"best-practics/common/consts"
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestInfof(t *testing.T) {
	InitZap()
	ctx := context.WithValue(context.Background(), consts.TraceKey, &common.Trace{TraceId: "123", Caller: "blog", UserId: 666})

	InfoF(ctx,"TEST_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
	ErrorF(ctx,"ERROR_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
}
