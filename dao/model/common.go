package model

import (
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}
