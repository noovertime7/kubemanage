package dao

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/source"
	"gorm.io/gorm"
	"strings"
)

func init() {
	source.RegisterInit(&SystemInitTable{})
}

type SystemInitTable struct {
	isInit bool
}

func (s *SystemInitTable) InitializerName() string {
	return strings.ToUpper("SystemInitTable")
}

func (s *SystemInitTable) MigrateTable(ctx context.Context, db *gorm.DB) error {
	tables := []interface{}{
		model.Workflow{},
		Admin{},
	}
	for _, t := range tables {
		if err := db.AutoMigrate(&t); err != nil {
			return err
		}
	}
	s.isInit = true
	return nil
}

func (s *SystemInitTable) InitializeData(ctx context.Context) error {
	return nil
}

func (s *SystemInitTable) TableCreated(ctx context.Context, db *gorm.DB) bool {
	tables := []interface{}{
		model.Workflow{},
		Admin{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}

func (s *SystemInitTable) DataInserted(ctx context.Context) bool {
	return s.isInit
}
