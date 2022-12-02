package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

// GetMenus
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      dto.Empty                                                  true  "空"
// @Success   200   {object}  middleware.Response{data=dto.SysMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /api//menu/get_menus [get]
func (m *menuController) GetMenus(ctx *gin.Context) {
	aid, err := utils.GetUserAuthorityId(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	menus, err := v1.CoreV1.Menu().GetMenu(ctx, aid)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, &dto.SysMenusResponse{Menus: menus})
}

// AddBaseMenu
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  middleware.Response{msg=string}  "新增菜单"
// @Router    /menu/add_base_menu [post]
func (m *menuController) AddBaseMenu(ctx *gin.Context) {
	params := &dto.AddSysMenusInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.Menu().AddBaseMenu(ctx, params); err != nil {
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
	if err := v1.CoreV1.Menu().AddMenuAuthority(ctx, params.Menus, params.AuthorityId); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}
