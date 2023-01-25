package cmdb

import (
	"context"
	"fmt"

	"github.com/noovertime7/kubemanage/runtime"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/common"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

type PermissionI interface {
	Save(ctx context.Context, search *model.Permission) error
	Updates(ctx context.Context, opt common.UpdateOption, search *model.Permission) error
	Find(ctx context.Context, search model.Permission) (*model.Permission, error)
	FindWithHosts(ctx context.Context, search model.Permission) (*model.Permission, error)
	FindList(ctx context.Context, search model.Permission) ([]*model.Permission, error)
	Delete(ctx context.Context, search model.Permission, isDelete bool) error

	PageList(ctx context.Context, params runtime.Pager) ([]*model.Permission, int64, error)
}

func NewPermissionI(db *gorm.DB) PermissionI {
	return &permission{db: db}
}

var _ PermissionI = &permission{}

type permission struct {
	db *gorm.DB
}

func (p *permission) Save(ctx context.Context, search *model.Permission) error {
	return p.db.WithContext(ctx).Create(&search).Error
}

func (p *permission) Updates(ctx context.Context, opt common.UpdateOption, search *model.Permission) error {
	query := opt(p.db)
	return query.WithContext(ctx).Updates(&search).Error
}

func (p *permission) Find(ctx context.Context, search model.Permission) (*model.Permission, error) {
	var out *model.Permission
	return out, p.db.WithContext(ctx).Where(&search).Find(&out).Error
}

func (p *permission) FindWithHosts(ctx context.Context, search model.Permission) (*model.Permission, error) {
	var out *model.Permission
	return out, p.db.WithContext(ctx).Preload("Hosts").Where(&search).Find(&out).Error
}

func (p *permission) FindList(ctx context.Context, search model.Permission) ([]*model.Permission, error) {
	var out []*model.Permission
	return out, p.db.WithContext(ctx).Preload("Hosts").Where(&search).Find(&out).Error
}

func (p *permission) Delete(ctx context.Context, search model.Permission, isDelete bool) error {
	if utils.IsStrEmpty(search.InstanceID) {
		return fmt.Errorf("instanceID is empty")
	}
	if isDelete {
		return p.db.WithContext(ctx).Where("instanceID = ?", search.InstanceID).Unscoped().Delete(&search).Error
	}
	return p.db.WithContext(ctx).Where("instanceID = ?", search.InstanceID).Delete(&search).Error
}

func (p *permission) PageList(ctx context.Context, params runtime.Pager) ([]*model.Permission, int64, error) {
	var total int64 = 0
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)
	query := p.db.WithContext(ctx).Where("")
	var list []*model.Permission
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
