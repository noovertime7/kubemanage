package dept

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"gorm.io/gorm"
)

type Department interface {
	Find(ctx context.Context, in *model.Department) (*model.Department, error)
	FindList(ctx context.Context, in *model.Department) ([]model.Department, error)
	FindListWithUsers(ctx context.Context, in *model.Department) (model.Department, error)
	Save(ctx context.Context, in *model.Department) error
	Updates(ctx context.Context, in *model.Department) error
}

type department struct {
	db *gorm.DB
}

var _ Department = &department{}

func NewDepartment(db *gorm.DB) Department {
	return &department{db: db}
}

func (d department) Find(ctx context.Context, in *model.Department) (*model.Department, error) {
	out := &model.Department{}
	return out, d.db.WithContext(ctx).Where(in).Find(&out).Error
}

func (d department) FindList(ctx context.Context, in *model.Department) ([]model.Department, error) {
	var out []model.Department
	return out, d.db.WithContext(ctx).Order("sort").Where(in).Find(&out).Error
}

func (d department) FindListWithUsers(ctx context.Context, in *model.Department) (model.Department, error) {
	var out model.Department
	return out, d.db.WithContext(ctx).Preload("SysUsers.Authorities").Order("sort").Where(in).Find(&out).Error
}

func (d department) Save(ctx context.Context, in *model.Department) error {
	return d.db.WithContext(ctx).Save(&in).Error
}

func (d department) Updates(ctx context.Context, in *model.Department) error {
	return d.db.WithContext(ctx).Updates(&in).Error
}
