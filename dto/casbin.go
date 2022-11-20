package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
)

// CasbinInfo Casbin info structure
type CasbinInfo struct {
	Path   string `form:"authorityId"  json:"path"`    // 路径
	Method string ` form:"authorityId"  json:"method"` // 方法
}

// CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint `form:"authorityId" json:"authorityId"` // 权限id
}

// BindingValidParams 绑定并校验参数
func (a *CasbinInReceive) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
