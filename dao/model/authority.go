package model

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func init() {
	RegisterInitializer(&SysAuthority{})
}

type SysAuthority struct {
	ID            int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	AuthorityId   int    `gorm:"column:authority_id;type:int(11) unsigned;AUTO_INCREMENT;comment:角色ID;primary_key" json:"authority_id"`
	AuthorityName string `gorm:"column:authority_name;type:varchar(191);comment:角色名" json:"authority_name"`
	DefaultRouter string `gorm:"column:default_router;type:varchar(191);default:dashboard;comment:默认菜单(dashboard)" json:"default_router"` // 父角色ID
	CommonModel
}

func (s *SysAuthority) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&s)
}

func (s *SysAuthority) InitData(ctx context.Context, db *gorm.DB) error {
	entities := []SysAuthority{
		{
			ID:            0,
			AuthorityId:   111,
			AuthorityName: "管理员",
			DefaultRouter: "dashboard",
			CommonModel: CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			ID:            0,
			AuthorityId:   222,
			AuthorityName: "普通用户",
			DefaultRouter: "dashboard",
			CommonModel: CommonModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}
	if err := db.WithContext(ctx).Create(&entities).Error; err != nil {
		auth := SysAuthority{}
		return errors.Wrap(err, auth.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (s *SysAuthority) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysAuthority
	if err := db.WithContext(ctx).Where("authority_id = '111' ").Find(&out).Error; err != nil {
		return false, err
	}
	return out.ID != 0, nil
}

func (s *SysAuthority) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&s)
}

func (*SysAuthority) TableName() string {
	return "sys_authorities"
}
