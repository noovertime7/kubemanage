package cmdb

import (
	"context"

	"github.com/noovertime7/kubemanage/dao/common"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
)

type HostGroup interface {
	Save(ctx context.Context, search model.CMDBHostGroup) error
	Find(ctx context.Context, search model.CMDBHostGroup) (model.CMDBHostGroup, error)
	FindList(ctx context.Context, search model.CMDBHostGroup) ([]model.CMDBHostGroup, error)
	FindListWithHosts(ctx context.Context, search model.CMDBHostGroup) ([]model.CMDBHostGroup, error)
	Updates(ctx context.Context, opt common.UpdateOption, search *model.CMDBHostGroup) error
	Delete(ctx context.Context, search model.CMDBHostGroup, isDelete bool) error
}

func NewHostGroup(db *gorm.DB) HostGroup {
	return &hostGroup{db: db}
}

var _ HostGroup = &hostGroup{}

type hostGroup struct {
	db *gorm.DB
}

func (h *hostGroup) Save(ctx context.Context, search model.CMDBHostGroup) error {
	return h.db.WithContext(ctx).Create(&search).Error
}

func (h *hostGroup) Find(ctx context.Context, search model.CMDBHostGroup) (model.CMDBHostGroup, error) {
	var out model.CMDBHostGroup
	return out, h.db.WithContext(ctx).Preload("Hosts").Where(&search).Find(&out).Error
}

func (h *hostGroup) Updates(ctx context.Context, opt common.UpdateOption, search *model.CMDBHostGroup) error {
	query := opt(h.db)
	return query.WithContext(ctx).Updates(&search).Error
}

func (h *hostGroup) FindList(ctx context.Context, search model.CMDBHostGroup) ([]model.CMDBHostGroup, error) {
	var out []model.CMDBHostGroup
	return out, h.db.WithContext(ctx).Where(&search).Order("sort desc").Find(&out).Error
}

func (h *hostGroup) FindListWithHosts(ctx context.Context, search model.CMDBHostGroup) ([]model.CMDBHostGroup, error) {
	var out []model.CMDBHostGroup
	return out, h.db.WithContext(ctx).Preload("Hosts").Where(&search).Order("sort desc").Find(&out).Error
}

func (h *hostGroup) Delete(ctx context.Context, search model.CMDBHostGroup, isDelete bool) error {
	if isDelete {
		return h.db.WithContext(ctx).Unscoped().Delete(&search).Error
	}
	return h.db.WithContext(ctx).Delete(&search).Error
}
