package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(WorkFlowOrder, &Workflow{})
}

type Workflow struct {
	ID          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Name        string `json:"name" gorm:"column:name"`
	NameSpace   string `json:"namespace" gorm:"column:namespace"`
	Replicas    int32  `json:"replicas" gorm:"column:replicas"`
	Deployment  string `json:"deployment" gorm:"column:deployment"`
	Service     string `json:"service" gorm:"column:service"`
	Ingress     string `json:"ingress" gorm:"column:ingress"`
	ServiceType string `json:"service_type" gorm:"column:service_type"`
	CommonModel
}

func (w *Workflow) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&w)
}

func (w *Workflow) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return true, nil
}

func (w *Workflow) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (w *Workflow) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(w)
}

func (w *Workflow) TableName() string {
	return "t_workflow"
}

func GetWorkflowTableName() string {
	temp := &Workflow{}
	return temp.TableName()
}
