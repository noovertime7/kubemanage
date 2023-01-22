package common

import "gorm.io/gorm"

// UpdateOption 用于将更新参数的能力给到service层
type UpdateOption func(db *gorm.DB) *gorm.DB
