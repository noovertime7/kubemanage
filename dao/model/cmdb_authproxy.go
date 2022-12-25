package model

import (
	"context"
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &cmdbAuthProxy{})
}

type cmdbAuthProxy struct {
	Id        uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserUUID  uuid.UUID
	Hosts     []cmdbHost `json:"hosts" gorm:"many2many:cmdb_proxy_host;"`
	StartTime time.Time
	EndTime   time.Time
	CommonModel
}

func (c *cmdbAuthProxy) TableName() string {
	return "cmdb_auth_proxy"
}

func (c *cmdbAuthProxy) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *cmdbAuthProxy) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *cmdbAuthProxy) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *cmdbAuthProxy) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
