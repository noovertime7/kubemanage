package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
)

type AddSysMenusInput struct {
	ParentId string `json:"parentId" comment:"父菜单ID" validate:"required"` // 父菜单ID
	Name     string `json:"name"  comment:"路由name" validate:"required"`   // 路由name
	Path     string `json:"path" comment:"路由path" validate:"required"`    // 路由path
	Disabled bool   `json:"disabled" comment:"是否禁用" validate:"required"`  // 是否在列表隐藏
	Hidden   bool   `json:"hidden" comment:"是否在列表隐藏" validate:"required"` // 是否在列表隐藏
	Sort     int    `json:"sort" comment:"排序标记" validate:"required"`      // 排序标记
	model.Meta
}

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []model.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenu `json:"menu"`
}

type AddMenuAuthorityInput struct {
	Menus       []model.SysBaseMenu `json:"menus"`
	AuthorityId uint                `json:"authorityId"  validate:"required"` // 角色ID
}

// BindingValidParams 绑定并校验参数
func (a *AddSysMenusInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

// BindingValidParams 绑定并校验参数
func (a *AddMenuAuthorityInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
