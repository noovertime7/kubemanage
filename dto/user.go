package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
	"gorm.io/gorm"
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

type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

type ChangeUserPwdInput struct {
	OldPwd string `json:"old_pwd" form:"old_pwd" comment:"原密码" validate:"required"`
	NewPwd string `json:"new_pwd" form:"new_pwd" comment:"new_pwd" validate:"required"`
}

// BindingValidParams 绑定并校验参数
func (a *ChangeUserPwdInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

type PageUsers struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page" form:"page"`         // 页码
	PageSize int             `json:"pageSize" form:"pageSize"` // 每页大小
	List     []model.SysUser `json:"list"`
}

type PageUsersIn struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

func (p PageUsersIn) GetPage() int {
	return p.Page
}

func (p PageUsersIn) GetPageSize() int {
	return p.PageSize
}

func (p PageUsersIn) IsFitter() bool {
	if p.UserName != "" || p.Phone != "" || p.Status != 0 {
		return true
	}
	return false
}

func (p PageUsersIn) Do(tx *gorm.DB) {
	if p.UserName != "" {
		tx.Where("user_name = ? ", p.UserName)
	}
	if p.Phone != "" {
		tx.Where("phone = ? ", p.Phone)
	}
	if p.Status != 0 {
		tx.Where("status = ? ", p.Status)
	}
}

func (p *PageUsersIn) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

// BindingValidParams 绑定并校验参数
func (a *SetUserAuth) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

// BindingValidParams 绑定并校验参数
func (a *AdminLoginInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
