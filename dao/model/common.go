package model

import (
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	ID        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
