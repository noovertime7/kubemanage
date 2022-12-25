package model

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(CMDBInitOrder, &cmdbHostGroup{})
}

type cmdbHostGroup struct {
	Id        uint   `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	ParentId  string `json:"parentId" gorm:"size:10"` //上层主机组
	GroupName string
	Sort      int             `json:"sort" gorm:"size:4;"` //排序
	Children  []cmdbHostGroup `json:"children" gorm:"-"`
	Hosts     []cmdbHost
	CommonModel
}

func (c *cmdbHostGroup) TableName() string {
	return "cmdb_host_group"
}

func (c *cmdbHostGroup) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&c)
}

func (c *cmdbHostGroup) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := c.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	return db.WithContext(ctx).Create(CMDBHostGroupInitData).Error
}

func (c *cmdbHostGroup) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out cmdbHostGroup
	if errors.Is(gorm.ErrRecordNotFound, db.WithContext(ctx).Where("group_name = 'Default'").
		First(&out).Error) { // 判断是否存在数据
		return false, nil
	}
	return true, nil
}

func (c *cmdbHostGroup) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&c)
}
