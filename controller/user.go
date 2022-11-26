package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/wonderivan/logger"
	"strconv"
)

type UserController struct{}

func UserRegister(group *gin.RouterGroup) {
	user := &UserController{}
	group.POST("/login", user.Login)
	group.GET("/loginout", user.LoginOut)
	group.GET("/getinfo", user.GetUserInfo)
	group.PUT("/:id/set_auth", user.SetUserAuthority)
	group.DELETE("/:id/delete_user", user.DeleteUser)
	group.POST("/:id/change_pwd", user.ChangePassword)
	group.PUT("/:id/reset_pwd", user.ResetPassword)
}

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/login [post]
func (u *UserController) Login(ctx *gin.Context) {
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

// LoginOut godoc
// @Summary 管理员退出登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/loginout
// @Accept  json
// @Produce  json
// @Param polygon body  true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/loginout [get]
func (u *UserController) LoginOut(ctx *gin.Context) {
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

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  middleware.Response{data=model.SysUser,msg=string}  "获取用户信息"
// @Router    /api/user/getinfo [get]
func (u *UserController) GetUserInfo(ctx *gin.Context) {
	clalms, err := utils.GetClaims(ctx)
	if err != nil {
		logger.Error("获取CustomClaims失败", err)
		return
	}
	userInfo, err := v1.CoreV1.User().GetUserInfo(ctx, clalms.ID, clalms.AuthorityId)
	if err != nil {
		logger.Error("获取userInfo失败", err)
		middleware.ResponseError(ctx, 30001, err)
	}
	middleware.ResponseSuccess(ctx, userInfo)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.SetUserAuth          true  "角色ID"
// @Success   200   {object}  middleware.Response{msg=string}  "设置用户权限"
// @Router    /api/user/{id}/set_auth [put]
func (u *UserController) SetUserAuthority(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
	}
	params := &dto.SetUserAuth{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := v1.CoreV1.User().SetUserAuth(ctx, uid, params.AuthorityId); err != nil {
		logger.Error("修改用户角色失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	// token中存在角色信息，需要生成新的token
	claims := utils.GetUserInfo(ctx)
	claims.AuthorityId = params.AuthorityId
	newToken, err := pkg.JWTToken.GenerateToken(claims.BaseClaims)
	if err != nil {
		logger.Error("修改用户角色失败,生成新token失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	ctx.Header("new-token", newToken)
	ctx.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
	middleware.ResponseSuccess(ctx, "操作成功")
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body                   true  "用户ID"
// @Success   200   {object}  middleware.Response{msg=string}  "删除用户"
// @Router    /api/user/{id}/delete_user [delete]
func (u *UserController) DeleteUser(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
	}
	if err := v1.CoreV1.User().DeleteUser(ctx, uid); err != nil {
		logger.Error("删除用户失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "操作成功")
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      dto.ChangeUserPwdInput    true  "用户ID, 原密码, 新密码"
// @Success   200   {object}  middleware.Response{msg=string}  "用户修改密码"
// @Router    /api/user/{id}/change_pwd [post]
func (u *UserController) ChangePassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
	}
	params := &dto.ChangeUserPwdInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := v1.CoreV1.User().ChangePassword(ctx, uid, params); err != nil {
		logger.Error("修改用户密码失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "操作成功")
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body                   true  "ID"
// @Success   200   {object}  middleware.Response{msg=string}  "重置用户密码"
// @Router    /api/user/{id}/reset_pwd [post]
func (u *UserController) ResetPassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
	}
	if err := v1.CoreV1.User().ResetPassword(ctx, uid); err != nil {
		logger.Error("重置用户密码失败", err.Error())
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	middleware.ResponseSuccess(ctx, "操作成功")
}
