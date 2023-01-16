package common

import "gorm.io/gorm"

// Option 用于将更新参数的能力给到service层
type Option func(db *gorm.DB) *gorm.DB
