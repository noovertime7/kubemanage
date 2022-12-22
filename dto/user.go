package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

type UserInfoInput struct {
	UserName     string `json:"userName" validate:"required" comment:"用户名"` // 用户登录名
	Password     string `json:"password"  validate:"" comment:"密码"`         // 用户登录密码
	NickName     string `json:"nickName" validate:"required" comment:"昵称"`  // 用户昵称 	// 活跃颜色
	DepartmentID uint   `json:"departmentId"  validate:"required"  comment:"部门ID"`
	AuthorityId  uint   `json:"authorityId" validate:"required" comment:"用户权限ID"` // 用户角色ID
	Authorities  []uint `json:"authorityIds"`
	Phone        string `json:"phone" `  // 用户手机号
	Email        string `json:"email"  ` // 用户邮箱
	Enable       int    `json:"enable" ` //用户是否被冻结 1正常 2冻结
}

// BindingValidParams 绑定并校验参数
func (a *UserInfoInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

type UserInfoOut struct {
	User      model.SysUser   `json:"user"`
	Menus     []model.SysMenu `json:"menus"`
	RuleNames []string        `json:"ruleNames"`
}

type SetUserAuthoritiesInput struct {
	Authorities []uint `json:"authorityIds"`
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
	Total    int64          `json:"total"`
	Page     int            `json:"page" form:"page"`         // 页码
	PageSize int            `json:"pageSize" form:"pageSize"` // 每页大小
	List     []PageUserItem `json:"list"`
}

type PageUserItem struct {
	ID             int                  `json:"id"`
	DepartmentID   uint                 `json:"departmentId" `
	DepartmentName string               `json:"departmentName"`
	UserName       string               `json:"userName"`  // 用户登录名 	// 用户登录密码
	NickName       string               `json:"nickName" ` // 用户昵称
	Authorities    []model.SysAuthority `json:"authorities"`
	Phone          string               `json:"phone" `  // 用户手机号
	Email          string               `json:"email"`   // 用户邮箱
	Enable         int                  `json:"enable" ` //用户是否被冻结 1正常 2冻结
	Status         int64                `json:"status"`
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
func (a *SetUserAuthoritiesInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

// BindingValidParams 绑定并校验参数
func (a *AdminLoginInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
