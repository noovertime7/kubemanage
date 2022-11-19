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

var InitializerList []*OrderedInitializer

// OrderedInitializer 组合一个顺序字段，以供排序
type OrderedInitializer struct {
	Order int
	DBInitializer
}

func RegisterInitializer(order int, c DBInitializer) {
	InitializerList = append(InitializerList, &OrderedInitializer{
		Order:         order,
		DBInitializer: c,
	})
	InitializerList = bubbleSort(InitializerList)
}

// 冒泡排序，根据order选择初始化顺序
func bubbleSort(s []*OrderedInitializer) []*OrderedInitializer {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j].Order > s[j+1].Order {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
