package authority

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"gorm.io/gorm"
)

type Authority interface {
	Find(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error)
	FindAllInfo(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error)
	FindList(ctx context.Context, authInfo *model.SysAuthority) ([]*model.SysAuthority, error)
	Save(ctx context.Context, authInfo *model.SysAuthority) error
	Updates(ctx context.Context, authInfo *model.SysAuthority) error
	Delete(ctx context.Context, authInfo *model.SysAuthority) error

	SetMenuAuthority(ctx context.Context, authInfo *model.SysAuthority) error
	//DeleteAuthorityMenu 解除权限与菜单的绑定关系
	DeleteAuthorityMenu(ctx context.Context, authInfo *model.SysAuthority, menus []model.SysBaseMenu) error
	PageList(ctx context.Context, params dto.PageInfo) ([]model.SysAuthority, int64, error)
}

var _ Authority = &authority{}

type authority struct {
	db *gorm.DB
}

func NewAuthority(db *gorm.DB) Authority {
	return &authority{db: db}
}

func (a *authority) Find(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error) {
	var out *model.SysAuthority
	return out, a.db.WithContext(ctx).Where(authInfo).Find(&out).Error
}

func (a *authority) FindAllInfo(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error) {
	var out *model.SysAuthority
	return out, a.db.WithContext(ctx).Preload("Users").Preload("SysBaseMenus").Where(authInfo).Find(&out).Error
}

func (a *authority) FindList(ctx context.Context, authInfo *model.SysAuthority) ([]*model.SysAuthority, error) {
	var out []*model.SysAuthority
	return out, a.db.WithContext(ctx).Where(&authInfo).Find(&out).Error
}

func (a *authority) PageList(ctx context.Context, params dto.PageInfo) ([]model.SysAuthority, int64, error) {
	var total int64 = 0
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	query := a.db.WithContext(ctx)
	var list []model.SysAuthority
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.Keyword != "" {
		query = query.Where("authority_name = ?", params.Keyword)
	}
	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (a *authority) Save(ctx context.Context, authInfo *model.SysAuthority) error {
	return a.db.WithContext(ctx).Create(&authInfo).Error
}

func (a *authority) Delete(ctx context.Context, authInfo *model.SysAuthority) error {
	return a.db.WithContext(ctx).Unscoped().Delete(authInfo).Error
}

func (a *authority) DeleteAuthorityMenu(ctx context.Context, authInfo *model.SysAuthority, menus []model.SysBaseMenu) error {
	return a.db.WithContext(ctx).Model(authInfo).Association("SysBaseMenus").Delete(menus)
}

func (a *authority) Updates(ctx context.Context, authInfo *model.SysAuthority) error {
	return a.db.WithContext(ctx).Updates(authInfo).Error
}

// SetMenuAuthority 菜单与角色绑定
func (a *authority) SetMenuAuthority(ctx context.Context, authInfo *model.SysAuthority) error {
	var s model.SysAuthority
	a.db.WithContext(ctx).Preload("SysBaseMenus").First(&s, "authority_id = ?", authInfo.AuthorityId)
	return a.db.WithContext(ctx).Model(&s).Association("SysBaseMenus").Replace(&authInfo.SysBaseMenus)
}
