package model

import (
	"context"

	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &CMDBSecret{})
}

type CMDBSecret struct {
	Id           uint   `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	InstanceID   string `json:"instanceID" gorm:"unique;not null;index;column:instanceID;comment:唯一id"`
	Name         string `json:"name" orm:"column:name;comment:名称"`
	Protocol     uint   `json:"protocol" gorm:"column:protocol;comment:协议1为ssh2为rdp"`
	SecretType   uint   `json:"secretType" gorm:"index;column:secretType;comment:认证类型1为密码2为秘钥"`
	HostUserName string `json:"hostUserName" gorm:"column:hostUserName;comment:主机用户名"`
	HostPassword string `json:"hostPassword" gorm:"column:hostPassword;comment:主机密码"`
	Content      string `json:"content" gorm:"column:content;comment:备注"`
	PrivateKey   string `json:"privateKey" gorm:"column:privateKey;comment:主机私钥"`
	CommonModel
}

func (c *CMDBSecret) TableName() string {
	return "cmdb_secret"
}

func (c *CMDBSecret) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *CMDBSecret) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *CMDBSecret) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *CMDBSecret) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
