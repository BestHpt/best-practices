/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package common

import (
	"best-practics/common/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper        *viper.Viper
	GlobalConfig config.GlobalConfig
	GVA_DB     *gorm.DB
)
