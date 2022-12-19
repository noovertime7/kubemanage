package model

import (
	"context"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/pkg"
)

func init() {
	RegisterInitializer(MenuAuthorityOrder, &MenuAuthority{})
}

type MenuAuthority struct{}

func (i MenuAuthority) TableName() string {
	return "sys_menu_authorities"
}

func (i *MenuAuthority) MigrateTable(ctx context.Context, db *gorm.DB) error {
	//数据由gorm填充不需要手动迁移
	return nil
}

func (i *MenuAuthority) InitData(ctx context.Context, db *gorm.DB) error {
	var (
		adminRole   = SysAuthorityEntities[0]
		userRole    = SysAuthorityEntities[1]
		userSubRole = SysAuthorityEntities[2]
		err         error
		ok          bool
	)
	ok, err = i.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	// admin 拥有全部菜单权限
	if err = db.Model(&adminRole).Association("SysBaseMenus").Append(SysBaseMenuEntities); err != nil {
		return err
	}

	// userRole cmdb菜单
	menu8881 := SysBaseMenuEntities[5:7]
	menu8881 = append(menu8881, SysBaseMenuEntities[0])
	menu8881 = append(menu8881, SysBaseMenuEntities[1])
	if err = db.Model(&userRole).Association("SysBaseMenus").Replace(menu8881); err != nil {
		return err
	}

	// userSubRole
	if err = db.Model(&userSubRole).Association("SysBaseMenus").Replace(SysBaseMenuEntities[:6]); err != nil {
		return err
	}
	if err = db.Model(&userSubRole).Association("SysBaseMenus").Append(SysBaseMenuEntities[5:6]); err != nil {
		return err
	}
	return nil
}

func (i *MenuAuthority) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	auth := &SysAuthority{}
	if err := db.WithContext(ctx).Model(auth).Where("authority_id = ?", pkg.AdminDefaultAuth).Preload("SysBaseMenus").Find(auth).Error; err != nil {
		return false, err
	}
	return len(auth.SysBaseMenus) > 0, nil
}

func (i *MenuAuthority) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return false // always replace
}
