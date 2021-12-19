package common

import (
	"time"
)

type BaseModel struct {
	ID        uint64       `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
