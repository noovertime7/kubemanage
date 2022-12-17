package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(SysApisInitOrder, &SysApi{})
}

type SysApi struct {
	ID          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Path        string `json:"path" gorm:"comment:api路径"`             // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`          // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	CommonModel
}

func (a *SysApi) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(a)
}

func (a *SysApi) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := a.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	return db.WithContext(ctx).Create(SysApis).Error
}

func (a *SysApi) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysApi
	// TODO 验证方式统一优化
	if err := db.WithContext(ctx).Where("path = '/api/user/login' ").Find(&out).Error; err != nil {
		return false, nil
	}
	return out.ID != 0, nil
}

func (a *SysApi) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(a)
}

func (*SysApi) TableName() string {
	return "sys_apis"
}
