package user

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/utils"

	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

func (u *userController) GetDepartmentTree(ctx *gin.Context) {
	data, err := v1.CoreV1.System().Department().GetDepartmentTree(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (u *userController) GetDepartmentUsers(ctx *gin.Context) {
	did, err := utils.ParseUint(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	data, err := v1.CoreV1.System().Department().GetDeptUsers(ctx, did)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
