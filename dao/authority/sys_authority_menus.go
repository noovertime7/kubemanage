package authority

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"gorm.io/gorm"
)

type AuthorityMenu interface {
	FindList(ctx context.Context, in *model.SysAuthorityMenu) ([]*model.SysAuthorityMenu, error)
}

var _ AuthorityMenu = &authorityMenu{}

type authorityMenu struct {
	db *gorm.DB
}

func NewAuthorityMenu(db *gorm.DB) *authorityMenu {
	return &authorityMenu{db: db}
}

func (a authorityMenu) FindList(ctx context.Context, in *model.SysAuthorityMenu) ([]*model.SysAuthorityMenu, error) {
	var out []*model.SysAuthorityMenu
	return out, a.db.WithContext(ctx).Where(&in).Find(&out).Error
}
