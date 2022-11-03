package model

import (
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
