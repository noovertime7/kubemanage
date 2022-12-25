package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &cmdbSecret{})
}

type cmdbSecret struct {
	Id           uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	HostUserName string
	HostPassword string
	PrivateKey   string
	CommonModel
}

func (c *cmdbSecret) TableName() string {
	return "cmdb_secret"
}

func (c *cmdbSecret) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *cmdbSecret) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *cmdbSecret) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *cmdbSecret) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
