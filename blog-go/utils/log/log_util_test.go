/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package log

import (
	"best-practics/common"
	"best-practics/common/consts"
	"best-practics/common/initialize"
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestInfof(t *testing.T) {
	common.Viper = initialize.InitViper("../../conf/config.yaml") // 初始化Viper
	InitZap()
	ctx := context.WithValue(context.Background(), consts.TraceKey, &common.Trace{TraceId: "123", Caller: "blog", UserId: 666})

	InfoF(ctx,"TEST_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
	ErrorF(ctx,"ERROR_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
}
