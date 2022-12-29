package model

import (
	"context"
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &CMDBAuthProxy{})
}

type CMDBAuthProxy struct {
	Id         uint      `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	InstanceID int64     `json:"instanceID" gorm:"index;column:instanceID;comment:唯一id"`
	UserUUID   uuid.UUID `json:"userUUID" gorm:"column:userUUID;comment:用户UUID"`
	UserName   string    `json:"userName" gorm:"column:userName;comment:用户名"`
	StartTime  time.Time `json:"startTime" gorm:"column:startTime;comment:开始时间"`
	EndTime    time.Time `json:"endTime" gorm:"column:endTime;comment:截止时间"`
	Hosts      []CMDBHost
	CommonModel
}

func (c *CMDBAuthProxy) TableName() string {
	return "cmdb_auth_proxy"
}

func (c *CMDBAuthProxy) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *CMDBAuthProxy) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *CMDBAuthProxy) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *CMDBAuthProxy) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
