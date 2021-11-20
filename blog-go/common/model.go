/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package common

import (
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
}
