package cmdb

import (
	"context"

	"github.com/noovertime7/kubemanage/dao/common"

	"github.com/noovertime7/kubemanage/runtime"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
)

type HostI interface {
	Save(ctx context.Context, search *model.CMDBHost) error
	Updates(ctx context.Context, opt common.UpdateOption, search *model.CMDBHost) error
	Find(ctx context.Context, search model.CMDBHost) (model.CMDBHost, error)
	FindList(ctx context.Context, search model.CMDBHost) ([]*model.CMDBHost, error)
	Delete(ctx context.Context, search model.CMDBHost, isDelete bool) error

	PageList(ctx context.Context, groupID uint, params runtime.Pager) ([]*model.CMDBHost, int64, error)
}

func NewHost(db *gorm.DB) HostI {
	return &host{db: db}
}

var _ HostI = &host{}

type host struct {
	db *gorm.DB
}

func (h *host) Save(ctx context.Context, search *model.CMDBHost) error {
	return h.db.WithContext(ctx).Create(&search).Error
}

func (h *host) Updates(ctx context.Context, opt common.UpdateOption, search *model.CMDBHost) error {
	query := opt(h.db)
	return query.WithContext(ctx).Updates(&search).Error
}

func (h *host) Find(ctx context.Context, search model.CMDBHost) (model.CMDBHost, error) {
	var out model.CMDBHost
	return out, h.db.WithContext(ctx).Where(&search).First(&out).Error
}

func (h *host) FindList(ctx context.Context, search model.CMDBHost) ([]*model.CMDBHost, error) {
	var out []*model.CMDBHost
	return out, h.db.WithContext(ctx).Where(&search).Find(&out).Error
}

func (h *host) Delete(ctx context.Context, search model.CMDBHost, isDelete bool) error {
	if isDelete {
		return h.db.WithContext(ctx).Where("instanceID = ?", search.InstanceID).Unscoped().Delete(&search).Error
	}
	return h.db.WithContext(ctx).Where("instanceID = ?", search.InstanceID).Delete(&search).Error
}

func (h *host) PageList(ctx context.Context, groupID uint, params runtime.Pager) ([]*model.CMDBHost, int64, error) {
	var total int64 = 0
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)
	query := h.db.WithContext(ctx).Where("cmdbHostGroupID = ?", groupID)
	var list []*model.CMDBHost
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.IsFitter() {
		params.Do(query)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id desc").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
