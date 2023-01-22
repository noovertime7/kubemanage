package model

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &Permission{})
}

type Permission struct {
	Id         uint       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	InstanceID string     `json:"instanceID" gorm:"index;column:instanceID;comment:唯一id"`
	UserUUID   uuid.UUID  `json:"userUUID" gorm:"column:userUUID;comment:用户UUID"`
	UserName   string     `json:"userName"`
	GroupName  string     `json:"groupName"`
	StartTime  time.Time  `json:"startTime" gorm:"column:startTime;comment:开始时间"`
	EndTime    time.Time  `json:"endTime" gorm:"column:endTime;comment:截止时间"`
	Hosts      []CMDBHost `json:"hosts"`
	Content    string     `json:"content" gorm:"column:content;comment:备注"`
	CommonModel
}

func (c *Permission) TableName() string {
	return "cmdb_permission"
}

func (c *Permission) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *Permission) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *Permission) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *Permission) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
