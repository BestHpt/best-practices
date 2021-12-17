package log

import (
	"best-practics/common"
	"best-practics/common/consts"
	"best-practics/common/initialize/viper"
	"best-practics/common/trace"
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestInfof(t *testing.T) {
	//1、初始化Viper
	common.Viper = viper.Init("../../../conf/config.yaml")
	InitZap()
	ctx := context.WithValue(context.Background(), consts.TraceKey, &trace.Trace{TraceId: "123", Caller: "blog", UserId: 666})

	InfoF(ctx,"TEST_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
	ErrorF(ctx,"ERROR_TAG",zap.Int("int_value", 47), zap.String("string_value", "48"))
}
