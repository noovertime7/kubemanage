package menu

import (
	"context"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
)

type BaseMenu interface {
	Find(ctx context.Context, in *model.SysBaseMenu) (*model.SysBaseMenu, error)
	FindIn(ctx context.Context, in []string) ([]*model.SysBaseMenu, error)
	FindList(ctx context.Context, in *model.SysBaseMenu) ([]model.SysBaseMenu, error)
	Save(ctx context.Context, in *model.SysBaseMenu) error
	Updates(ctx context.Context, in *model.SysBaseMenu) error
}

type baseMenu struct {
	db *gorm.DB
}

func NewBaseMenu(db *gorm.DB) *baseMenu {
	return &baseMenu{db: db}
}

func (b baseMenu) Find(ctx context.Context, in *model.SysBaseMenu) (*model.SysBaseMenu, error) {
	var out *model.SysBaseMenu
	return out, b.db.WithContext(ctx).Where(in).Find(&out).Error
}

func (b baseMenu) FindIn(ctx context.Context, in []string) ([]*model.SysBaseMenu, error) {
	//做一下排序
	var out []*model.SysBaseMenu
	return out, b.db.WithContext(ctx).Where("id in (?)", in).Order("sort").Find(&out).Error
}

func (b baseMenu) FindList(ctx context.Context, in *model.SysBaseMenu) ([]model.SysBaseMenu, error) {
	var out []model.SysBaseMenu
	return out, b.db.WithContext(ctx).Order("sort").Where(in).Find(&out).Error
}

func (b baseMenu) Save(ctx context.Context, in *model.SysBaseMenu) error {
	return b.db.WithContext(ctx).Create(in).Error
}

func (b baseMenu) Updates(ctx context.Context, in *model.SysBaseMenu) error {
	return b.db.WithContext(ctx).Updates(in).Error
}
