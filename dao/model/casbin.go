package model

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	adapter "github.com/casbin/gorm-adapter/v3"
)

func init() {
	RegisterInitializer(CasbinInitOrder, &casbinDateBase{})
}

type casbinDateBase struct{}

func (c casbinDateBase) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c casbinDateBase) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&adapter.CasbinRule{})
}

func (c casbinDateBase) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := c.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	return db.WithContext(ctx).Create(CasbinApi).Error
}

func (c casbinDateBase) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	// TODO 管理员用户判断登陆接口是否具有权限
	if errors.Is(gorm.ErrRecordNotFound, db.WithContext(ctx).Where(adapter.CasbinRule{Ptype: "p", V0: "222", V1: "/api/user/login", V2: "POST"}).
		First(&adapter.CasbinRule{}).Error) { // 判断是否存在数据
		return false, nil
	}
	return true, nil
}

func (c casbinDateBase) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&adapter.CasbinRule{})
}
