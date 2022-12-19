package user

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"strconv"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

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
func (u *userController) Login(ctx *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	token, err := v1.CoreV1.System().User().Login(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.LoginErr, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.LoginErr, err))
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
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/loginout [get]
func (u *userController) LoginOut(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		v1.Log.Error(globalError.ServerError)
	}
	cla, _ := claims.(*pkg.CustomClaims)
	if err := v1.CoreV1.System().User().LoginOut(ctx, cla.ID); err != nil {
		v1.Log.ErrorWithCode(globalError.LogoutErr, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
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
func (u *userController) GetUserInfo(ctx *gin.Context) {
	clalms, err := utils.GetClaims(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		return
	}
	userInfo, err := v1.CoreV1.System().User().GetUserInfo(ctx, clalms.ID, clalms.AuthorityId)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.LogoutErr, err))
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
func (u *userController) SetUserAuthority(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	params := &dto.SetUserAuth{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().User().SetUserAuth(ctx, uid, params.AuthorityId); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	// token中存在角色信息，需要生成新的token
	claims := utils.GetUserInfo(ctx)
	claims.AuthorityId = params.AuthorityId
	newToken, err := pkg.JWTToken.GenerateToken(claims.BaseClaims)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
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
// @Success   200   {object}  middleware.Response{msg=string}  "删除用户"
// @Router    /api/user/{id}/delete_user [delete]
func (u *userController) DeleteUser(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.System().User().DeleteUser(ctx, uid); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
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
func (u *userController) ChangePassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	params := &dto.ChangeUserPwdInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().User().ChangePassword(ctx, uid, params); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Success   200   {object}  middleware.Response{msg=string}  "重置用户密码"
// @Router    /api/user/{id}/reset_pwd [put]
func (u *userController) ResetPassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().User().ResetPassword(ctx, uid); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "操作成功")
}

func (u *userController) PageUsers(ctx *gin.Context) {
	did, err := utils.ParseUint(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	params := &dto.PageUsersIn{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.System().User().PageList(ctx, did, *params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
