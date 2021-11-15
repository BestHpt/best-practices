package global

import (
	"best-practics/common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Viper        *viper.Viper
	Logger       *zap.Logger
	GlobalConfig config.GlobalConfig
)
