package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

type MenuController struct{}

func MenuRegister(group *gin.RouterGroup) {
	menu := &MenuController{}
	group.GET("/get_menus", menu.GetMenus)
}

func (m *MenuController) GetMenus(ctx *gin.Context) {
	aid, err := utils.GetUserAuthorityId(ctx)
	if err != nil {
		logger.Error("获取菜单失败", err.Error())
		middleware.ResponseError(ctx, 10001, err)
		return
	}
	menus, err := v1.CoreV1.Menu().GetMenu(ctx, aid)
	if err != nil {
		logger.Error("获取菜单失败", err.Error())
		middleware.ResponseError(ctx, 10002, err)
		return
	}
	middleware.ResponseSuccess(ctx, &dto.SysMenusResponse{Menus: menus})
}
