package model

import (
	"gorm.io/gorm"
	"time"
)

type GormModel struct {
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeleteAt  *gorm.DeletedAt `json:"delete_at"`
}
