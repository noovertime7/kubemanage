package dept

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/runtime"
	"gorm.io/gorm"
)

type Department interface {
	Find(ctx context.Context, in *model.Department) (*model.Department, error)
	FindList(ctx context.Context, in *model.Department) ([]model.Department, error)
	FindListWithUsers(ctx context.Context, in *model.Department) (model.Department, error)
	Save(ctx context.Context, in *model.Department) error
	Updates(ctx context.Context, in *model.Department) error

	PageList(ctx context.Context, params runtime.Pager) ([]model.Department, int64, error)
}

type department struct {
	db *gorm.DB
}

var _ Department = &department{}

func NewDepartment(db *gorm.DB) Department {
	return &department{db: db}
}

func (d *department) Find(ctx context.Context, in *model.Department) (*model.Department, error) {
	out := &model.Department{}
	return out, d.db.WithContext(ctx).Where(in).Find(&out).Error
}

func (d *department) FindList(ctx context.Context, in *model.Department) ([]model.Department, error) {
	var out []model.Department
	return out, d.db.WithContext(ctx).Order("sort").Where(in).Find(&out).Error
}

func (d *department) FindListWithUsers(ctx context.Context, in *model.Department) (model.Department, error) {
	var out model.Department
	return out, d.db.WithContext(ctx).Preload("SysUsers.Authorities").Order("sort").Where(in).Find(&out).Error
}

func (d *department) Save(ctx context.Context, in *model.Department) error {
	return d.db.WithContext(ctx).Save(&in).Error
}

func (d *department) Updates(ctx context.Context, in *model.Department) error {
	return d.db.WithContext(ctx).Updates(&in).Error
}

func (d *department) PageList(ctx context.Context, params runtime.Pager) ([]model.Department, int64, error) {
	var total int64 = 0
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)
	query := d.db.WithContext(ctx).Where("")
	var list []model.Department
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.IsFitter() {
		params.Do(query)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
