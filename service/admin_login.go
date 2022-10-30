package service

import (
	"github.com/noovertime7/kubemanage/dao"
	dtouc "github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg"
)

var Admin admin

type admin struct {
	UserName     string
	Password     string
	SaltPassword string
}

func (a *admin) Login(loginInfo *dtouc.AdminLoginInput) (string, error) {
	var admindb = &dao.Admin{}
	//检查用户名密码
	userinfo, err := admindb.LoginCheck(loginInfo)
	if err != nil {
		return "", err
	}
	//生成token
	token, err := pkg.GenerateToken(&userinfo.Id)
	if err != nil {
		return "", err
	}
	//更改在线状态
	userinfo.Status = 1
	if err := userinfo.UpdateStatus(); err != nil {
		return "", err
	}
	return token, err
}

func (a *admin) Logout(uid int) error {
	AdminDB := &dao.Admin{Id: uid, Status: 0}
	return AdminDB.UpdateStatus()
}
