package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
)

// CasbinInfo Casbin info structure
type CasbinInfo struct {
	Path   string `form:"path"  json:"path"`      // 路径
	Method string ` form:"method"  json:"method"` // 方法
}

func (b CasbinInfo) GetPATH() string {
	return b.Path
}

func (b CasbinInfo) GetMethod() string {
	return b.Method
}

// UpdateCasbinInput 通过角色id更改接口权限
type UpdateCasbinInput struct {
	AuthorityId uint         `form:"authorityId" json:"authorityId"` // 权限id
	CasbinInfo  []CasbinInfo `json:"casbinInfos"`
}

// CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint `form:"authorityId" json:"authorityId"` // 权限id
}

// BindingValidParams 绑定并校验参数
func (a *CasbinInReceive) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

// BindingValidParams 绑定并校验参数
func (a *UpdateCasbinInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
