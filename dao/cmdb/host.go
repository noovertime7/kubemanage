package cmdb

import (
	"context"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
)

type HostI interface {
	Save(ctx context.Context, search model.CMDBHost) error
	Find(ctx context.Context, search model.CMDBHost) (model.CMDBHost, error)
	FindList(ctx context.Context, search model.CMDBHost) ([]*model.CMDBHost, error)
	Delete(ctx context.Context, search model.CMDBHost, isDelete bool) error
}

func NewHost(db *gorm.DB) HostI {
	return &host{db: db}
}

var _ HostI = &host{}

type host struct {
	db *gorm.DB
}

func (h *host) Save(ctx context.Context, search model.CMDBHost) error {
	return h.db.WithContext(ctx).Create(&search).Error
}

func (h *host) Find(ctx context.Context, search model.CMDBHost) (model.CMDBHost, error) {
	var out model.CMDBHost
	return out, h.db.WithContext(ctx).Where(&search).Find(&out).Error
}

func (h *host) FindList(ctx context.Context, search model.CMDBHost) ([]*model.CMDBHost, error) {
	var out []*model.CMDBHost
	return out, h.db.WithContext(ctx).Where(&search).Find(&out).Error
}

func (h *host) Delete(ctx context.Context, search model.CMDBHost, isDelete bool) error {
	if isDelete {
		return h.db.WithContext(ctx).Unscoped().Delete(&search).Error
	}
	return h.db.WithContext(ctx).Delete(&search).Error
}
