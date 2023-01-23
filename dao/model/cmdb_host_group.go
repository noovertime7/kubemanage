package model

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &CMDBHostGroup{})
}

type CMDBHostGroup struct {
	Id         uint            `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	InstanceID string          `json:"instanceID" gorm:"unique;not null;index;column:instanceID;comment:唯一id"`
	ParentId   string          `json:"parentId" gorm:"size:10;comment:主机组父ID"` //上层主机组
	GroupName  string          `json:"groupName" gorm:"index;column:groupName;comment:主机组名称"`
	Sort       uint            `json:"sort" gorm:"size:4;"` //排序
	Children   []CMDBHostGroup `json:"children" gorm:"-"`
	Hosts      []CMDBHost      `json:"hosts"`
	HostNum    int64           `json:"hostNum" gorm:"-"`
	CommonModel
}

func (c *CMDBHostGroup) TableName() string {
	return "cmdb_host_group"
}

func (c *CMDBHostGroup) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *CMDBHostGroup) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := c.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	return db.WithContext(ctx).Create(CMDBHostGroupInitData).Error
}

func (c *CMDBHostGroup) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out CMDBHostGroup
	if errors.Is(gorm.ErrRecordNotFound, db.WithContext(ctx).Where("groupName = 'Kubemanage'").
		First(&out).Error) { // 判断是否存在数据
		return false, nil
	}
	return true, nil
}

func (c *CMDBHostGroup) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
