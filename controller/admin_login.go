package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/wonderivan/logger"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/loginout", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /admin_login/login [post]
func (a *AdminLoginController) AdminLogin(ctx *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	token, err := v1.CoreV1.User().Login(ctx, params)
	if err != nil {
		logger.Error("登录失败", err.Error())
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, &dto.AdminLoginOut{Token: token})
}

func (a *AdminLoginController) AdminLoginOut(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		logger.Error("claims不存在,请检查jwt中间件")
	}
	cla, _ := claims.(*pkg.CustomClaims)
	if err := v1.CoreV1.User().LoginOut(ctx, cla.ID); err != nil {
		logger.Error("退出登录失败", err)
		middleware.ResponseSuccess(ctx, "退出失败")
		return
	}
	middleware.ResponseSuccess(ctx, "退出成功")
}
