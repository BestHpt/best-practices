/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package common

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper        *viper.Viper
	GVA_DB       *gorm.DB
)
