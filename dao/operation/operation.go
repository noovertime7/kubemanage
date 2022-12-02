package operation

import (
	"context"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
)

type Operation interface {
	Find(ctx context.Context, in *model.SysOperationRecord) (*model.SysOperationRecord, error)
	PageList(ctx context.Context, params *dto.OperationListInput) ([]*model.SysOperationRecord, int64, error)
	Save(ctx context.Context, in *model.SysOperationRecord) error
	Delete(ctx context.Context, in *model.SysOperationRecord) error
}

type operation struct {
	db *gorm.DB
}

func (o *operation) Find(ctx context.Context, in *model.SysOperationRecord) (*model.SysOperationRecord, error) {
	out := &model.SysOperationRecord{}
	return out, o.db.WithContext(ctx).Where(in).Find(&out).Error
}

func (o *operation) Save(ctx context.Context, in *model.SysOperationRecord) error {
	return o.db.WithContext(ctx).Create(in).Error
}

func (o *operation) Delete(ctx context.Context, in *model.SysOperationRecord) error {
	return o.db.WithContext(ctx).Delete(in).Error
}

func NewOperation(db *gorm.DB) *operation {
	return &operation{db: db}
}

func (o *operation) PageList(ctx context.Context, params *dto.OperationListInput) ([]*model.SysOperationRecord, int64, error) {
	var total int64 = 0
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	query := o.db.WithContext(ctx)
	var list []*model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.Method != "" {
		query = query.Where("method = ?", params.Method)
	}
	if params.Path != "" {
		query = query.Where("path = ?", params.Path)
	}
	if params.Status != 0 {
		query = query.Where("status = ?", params.Status)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
