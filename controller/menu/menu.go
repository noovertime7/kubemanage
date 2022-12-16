package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

// GetMenusByAuthID
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      dto.Empty                                                  true  "空"
// @Success   200   {object}  middleware.Response{data=dto.SysBaseMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /api/menu/:authID/getMenuByAuthID [get]
func (m *menuController) GetMenusByAuthID(ctx *gin.Context) {
	authID, err := utils.ParseUint(ctx.Param("authID"))
	if err != nil || authID == 0 {
		v1.Log.ErrorWithCode(globalError.ParamBindError, fmt.Errorf("authID empty"))
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, fmt.Errorf("authID empty")))
	}
	menus, err := v1.CoreV1.System().Menu().GetMenuByAuthorityID(ctx, authID)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, &dto.SysMenusResponse{Menus: menus})
}

// GetBaseMenus
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      dto.Empty                                                  true  "空"
// @Success   200   {object}  middleware.Response{data=dto.SysBaseMenusResponse,msg=string}  "获取系统菜单详情列表"
// @Router    /api/menu/getBaseMenuTree [get]
func (m *menuController) GetBaseMenus(ctx *gin.Context) {
	menus, err := v1.CoreV1.System().Menu().GetBassMenu(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, menus)
}

// AddBaseMenu
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  middleware.Response{msg=string}  "新增菜单"
// @Router    /api/menu/add_base_menu [post]
func (m *menuController) AddBaseMenu(ctx *gin.Context) {
	params := &dto.AddSysMenusInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.System().Menu().AddBaseMenu(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}

// AddMenuAuthority
// @Tags      AuthorityMenu
// @Summary   增加menu和角色关联关系
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AddMenuAuthorityInput  true  "角色ID"
// @Success   200   {object}  response.Response{msg=string}   "增加menu和角色关联关系"
// @Router    /api/menu/add_menu_authority [post]
func (m *menuController) AddMenuAuthority(ctx *gin.Context) {
	params := &dto.AddMenuAuthorityInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.System().Menu().AddMenuAuthority(ctx, params.Menus, params.AuthorityId); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}
