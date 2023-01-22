package cmdb

import (
	"context"

	"github.com/noovertime7/kubemanage/dao/common"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/runtime"
)

type SecretI interface {
	Save(ctx context.Context, search *model.CMDBSecret) error
	Updates(ctx context.Context, opt common.UpdateOption, in *model.CMDBSecret) error
	Find(ctx context.Context, search model.CMDBSecret) (model.CMDBSecret, error)
	FindList(ctx context.Context, search model.CMDBSecret) ([]model.CMDBSecret, error)
	Delete(ctx context.Context, search model.CMDBSecret, isDelete bool) error

	PageList(ctx context.Context, params runtime.Pager) ([]model.CMDBSecret, int64, error)
}

type secret struct {
	db *gorm.DB
}

func NewSecretI(db *gorm.DB) SecretI {
	return &secret{db: db}
}

func (s *secret) Save(ctx context.Context, search *model.CMDBSecret) error {
	return s.db.WithContext(ctx).Where(&search).Create(&search).Error
}

func (s *secret) Updates(ctx context.Context, opt common.UpdateOption, in *model.CMDBSecret) error {
	query := opt(s.db)
	return query.WithContext(ctx).Updates(&in).Error
}

func (s *secret) Find(ctx context.Context, search model.CMDBSecret) (model.CMDBSecret, error) {
	var out model.CMDBSecret
	return out, s.db.WithContext(ctx).Where(&search).First(&out).Error
}

func (s *secret) FindList(ctx context.Context, search model.CMDBSecret) ([]model.CMDBSecret, error) {
	var out []model.CMDBSecret
	return out, s.db.WithContext(ctx).Where(&search).Find(&out).Error
}

func (s *secret) Delete(ctx context.Context, search model.CMDBSecret, isDelete bool) error {
	if isDelete {
		return s.db.WithContext(ctx).Where(&search).Unscoped().Delete(&search).Error
	}
	return s.db.WithContext(ctx).Where(&search).Delete(&search).Error
}

func (s *secret) PageList(ctx context.Context, params runtime.Pager) ([]model.CMDBSecret, int64, error) {
	var total int64 = 0
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)
	query := s.db.WithContext(ctx).Where("")
	var list []model.CMDBSecret
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
