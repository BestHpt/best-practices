package global

import (
	"best-practics/common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Viper        *viper.Viper
	Logger       *zap.Logger
	GlobalConfig config.GlobalConfig
	GVA_DB     *gorm.DB
)
