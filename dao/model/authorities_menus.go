package model

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	)
	// admin
	if err = db.Model(&adminRole).Association("SysBaseMenus").Replace(SysBaseMenuEntities[:6]); err != nil {
		return err
	}
	if err = db.Model(&adminRole).Association("SysBaseMenus").Append(SysBaseMenuEntities[6:]); err != nil {
		return err
	}

	// userRole
	menu8881 := SysBaseMenuEntities[:6]
	menu8881 = append(menu8881, SysBaseMenuEntities[6])
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
	if ret := db.Model(auth).
		Where("authority_id = ?", 9528).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false, ret.Error
		}
		return len(auth.SysBaseMenus) > 0, nil
	}
	return false, errors.New("MenuAuthority IsInitData failed")
}

func (i *MenuAuthority) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return false // always replace
}
