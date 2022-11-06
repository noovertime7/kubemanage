package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/source"
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
		model.UserModel{},
	}
	for _, t := range tables {
		if err := db.AutoMigrate(&t); err != nil {
			return err
		}
	}
	s.isInit = true
	return nil
}

func (s *SystemInitTable) InitializeData(ctx context.Context, db *gorm.DB) error {
	datas := []*model.UserModel{{
		UserName: "admin",
		Salt:     "admin",
		Password: "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
		Status:   sql.NullInt64{Int64: 0, Valid: true},
		CommonModel: model.CommonModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}}
	for _, data := range datas {
		if err := db.Create(data).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *SystemInitTable) TableCreated(ctx context.Context, db *gorm.DB) bool {
	tables := []interface{}{
		model.Workflow{},
		model.UserModel{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}

func (s *SystemInitTable) DataInserted(ctx context.Context, db *gorm.DB) bool {
	tempUser := &model.UserModel{}
	if err := db.WithContext(ctx).Table("t_user").Where("user_name = 'admin' ").Find(tempUser).Error; err != nil {
		return false
	}
	return tempUser.ID != 0
}
