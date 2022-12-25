package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &cmdbHost{})
}

type cmdbHost struct {
	Id           uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Address      string
	Port         string
	HostUserName string
	HostPassword string
	PrivateKey   string
	SecretID     uint
	Status       uint
	AuthProxys   []cmdbAuthProxy `json:"authProxys" gorm:"many2many:cmdb_proxy_host;"`
	CommonModel
}

func (c *cmdbHost) TableName() string {
	return "cmdb_host"
}

func (c *cmdbHost) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *cmdbHost) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *cmdbHost) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *cmdbHost) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
