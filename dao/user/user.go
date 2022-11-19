package user

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User interface {
	Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error)
	Save(ctx context.Context, userInfo *model.SysUser) error
	Updates(ctx context.Context, userInfo *model.SysUser) error
}

var _ User = &user{}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error) {
	user := &model.SysUser{}
	if err := u.db.WithContext(ctx).Where(userInfo).First(user).Error; err != nil {
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

func (u *user) Updates(ctx context.Context, userInfo *model.SysUser) error {
	if userInfo.ID == 0 {
		return errors.New("id not set")
	}
	return u.db.WithContext(ctx).Updates(userInfo).Error
}
