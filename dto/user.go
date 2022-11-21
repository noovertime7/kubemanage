package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
)

type AdminLoginInput struct {
	UserName string `form:"username" json:"username" comment:"用户名"  validate:"required" example:"用户名"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"密码"`
}

type AdminLoginOut struct {
	Token string `form:"token" json:"token" comment:"token"  example:"token"`
}

type UserInfoOut struct {
	User      model.SysUser   `json:"user"`
	Menus     []model.SysMenu `json:"menus"`
	RuleNames []string        `json:"ruleNames"`
}

// BindingValidParams 绑定并校验参数
func (a *AdminLoginInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
