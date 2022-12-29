package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &CMDBHost{})
}

type CMDBHost struct {
	Id              uint   `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	InstanceID      int64  `json:"instanceID" gorm:"index;column:instanceID;comment:唯一id"`
	Address         string `json:"address" gorm:"column:address;comment:主机地址"`
	Port            string `json:"port" gorm:"column:port;comment:主机端口"`
	HostUserName    string `json:"hostUserName" gorm:"column:hostUserName;comment:主机用户名"`
	HostPassword    string `json:"hostPassword" gorm:"column:hostPassword;comment:主机密码"`
	PrivateKey      string `json:"privateKey" gorm:"column:privateKey;comment:主机私钥"`
	SecretID        uint   `json:"secretID" gorm:"column:secretID;comment:主机认证id"`
	Status          uint   `json:"status" gorm:"column:status;comment:主机状态"`
	CMDBAuthProxyID uint   `json:"cmdbAuthProxyID" gorm:"column:cmdbAuthProxyID;comment:主机策略id"`
	CMDBHostGroupID uint   `json:"cmdbHostGroupID" gorm:"column:cmdbHostGroupID;comment:主机组id"`
	CommonModel
}

func (c *CMDBHost) TableName() string {
	return "cmdb_host"
}

func (c *CMDBHost) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *CMDBHost) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (c *CMDBHost) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return false, nil
}

func (c *CMDBHost) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
