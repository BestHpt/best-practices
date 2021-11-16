package initialize

import (
	"best-practics/common/global"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	global.Viper = InitViper("../../conf/config.yaml")  // 初始化Viper
	global.Logger = InitZap()                       // 初始化zap日志库
	global.Logger.Debug("debug log", zap.Int("line", 47), zap.Int("line22", 48))
	global.Logger.Info("Info log", zap.Any("level", "1231231231"))
	global.Logger.Warn("warn log", zap.String("level", `{"a":"4","b":"5"}`))
	global.Logger.Error("err log", zap.String("level", `{"a":"7","b":"8"}`))
}
