package user

import (
	"context"
	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/runtime"
)

type User interface {
	Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error)
	Save(ctx context.Context, userInfo *model.SysUser) error
	Updates(ctx context.Context, userInfo *model.SysUser) error
	Delete(ctx context.Context, userInfo *model.SysUser) error
	PageList(ctx context.Context, did uint, params runtime.Pager) ([]model.SysUser, int64, error)
	// ReplaceAuthorities  绑定用户与角色
	ReplaceAuthorities(ctx context.Context, userInfo *model.SysUser, auths []model.SysAuthority) error
	// RemoveAuthorities  移除用户与角色
	RemoveAuthorities(ctx context.Context, userInfo *model.SysUser, auths []model.SysAuthority) error
}

var _ User = &user{}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) PageList(ctx context.Context, did uint, params runtime.Pager) ([]model.SysUser, int64, error) {
	var total int64 = 0
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)
	query := u.db.WithContext(ctx)
	var list []model.SysUser
	query = query.Where("department_id = ? ", did)
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.IsFitter() {
		params.Do(query)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("Authorities").Order("created_at").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (u *user) Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error) {
	user := &model.SysUser{}
	if err := u.db.WithContext(ctx).Preload("Authorities").Preload("Authority").Where(userInfo).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *user) Save(ctx context.Context, userInfo *model.SysUser) error {
	return u.db.WithContext(ctx).Save(userInfo).Error
}

// ReplaceAuthorities 给用户增加权限
func (u *user) ReplaceAuthorities(ctx context.Context, userInfo *model.SysUser, auths []model.SysAuthority) error {
	return u.db.WithContext(ctx).Model(&userInfo).Association("Authorities").Replace(&auths)
}

func (u *user) RemoveAuthorities(ctx context.Context, userInfo *model.SysUser, auths []model.SysAuthority) error {
	return u.db.WithContext(ctx).Model(&userInfo).Association("Authorities").Delete(&auths)
}

func (u *user) Updates(ctx context.Context, userInfo *model.SysUser) error {
	return u.db.WithContext(ctx).Updates(userInfo).Error
}

func (u *user) Delete(ctx context.Context, userInfo *model.SysUser) error {
	return u.db.WithContext(ctx).Delete(userInfo).Error
}
