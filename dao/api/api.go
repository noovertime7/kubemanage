package api

import (
	"context"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
)

type APi interface {
	FindList(ctx context.Context, search model.SysApi) ([]model.SysApi, error)
}

var _ APi = &api{}

func NewApi(db *gorm.DB) APi {
	return &api{db: db}
}

type api struct {
	db *gorm.DB
}

func (a *api) FindList(ctx context.Context, search model.SysApi) ([]model.SysApi, error) {
	var out []model.SysApi
	return out, a.db.WithContext(ctx).Where(&search).Find(&out).Error
}
