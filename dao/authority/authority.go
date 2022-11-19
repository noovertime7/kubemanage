package authority

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"gorm.io/gorm"
)

type Authority interface {
	Find(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error)
	FindList(ctx context.Context, authInfo *model.SysAuthority) ([]*model.SysAuthority, error)
	Save(ctx context.Context, authInfo *model.SysAuthority) error
	Updates(ctx context.Context, authInfo *model.SysAuthority) error
}

var _ Authority = &authority{}

type authority struct {
	db *gorm.DB
}

func NewAuthority(db *gorm.DB) *authority {
	return &authority{db: db}
}

func (a *authority) Find(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error) {
	var out *model.SysAuthority
	return out, a.db.WithContext(ctx).Where(authInfo).Find(out).Error
}

func (a *authority) FindList(ctx context.Context, authInfo *model.SysAuthority) ([]*model.SysAuthority, error) {
	var out []*model.SysAuthority
	return out, a.db.WithContext(ctx).Where(&authInfo).Find(&out).Error
}

func (a *authority) Save(ctx context.Context, authInfo *model.SysAuthority) error {
	return a.db.WithContext(ctx).Create(authInfo).Error
}

func (a *authority) Updates(ctx context.Context, authInfo *model.SysAuthority) error {
	return a.db.WithContext(ctx).Updates(authInfo).Error
}
