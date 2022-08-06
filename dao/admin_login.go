package dao

import (
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/public"
	"github.com/pkg/errors"
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	Status    int       `json:"status" gorm:"column:status" description:"登录状态"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "t_admin"
}

func (a *Admin) LoginCheck(param *dto.AdminLoginInput) (*Admin, error) {
	admininfo, err := a.Find(&Admin{UserName: param.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := public.GenSaltPassword(admininfo.Salt, param.Password)
	if admininfo.Password != saltPassword {
		return nil, errors.New("密码错误请重新输入")
	}
	return admininfo, nil
}

func (a *Admin) Find(search *Admin) (*Admin, error) {
	out := &Admin{}
	if err := Gorm.Where(search).Find(out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (a *Admin) Save() error {
	return Gorm.Save(a).Error
}

func (a *Admin) UpdateStatus() error {
	return Gorm.Table(a.TableName()).Where("id = ?", a.Id).Updates(map[string]interface{}{
		"status": a.Status,
	}).Error
}
