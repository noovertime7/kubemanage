package model

import (
	"context"
	"gorm.io/gorm"
)

// DBInitializer 用于初始化数据表与数据
type DBInitializer interface {
	TableName() string
	MigrateTable(ctx context.Context, db *gorm.DB) error
	InitData(ctx context.Context, db *gorm.DB) error
	IsInitData(ctx context.Context, db *gorm.DB) (bool, error)
	TableCreated(ctx context.Context, db *gorm.DB) bool
}

var InitializerList []DBInitializer

func RegisterInitializer(c DBInitializer) {
	InitializerList = append(InitializerList, c)
}
