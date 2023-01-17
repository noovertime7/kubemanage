package cmdb

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/utils"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

func (c *cmdbController) CreateHost(ctx *gin.Context) {
	params := &dto.CMDBHostCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().CreateHost(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func (c *cmdbController) UpdateHost(ctx *gin.Context) {
	params := &dto.CMDBHostCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().UpdateHost(ctx, params); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func (c *cmdbController) PageHost(ctx *gin.Context) {
	groupID, err := utils.ParseUint(ctx.Param("groupID"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	params := &dto.PageListCMDBHostInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.CMDB().Host().PageHost(ctx, groupID, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (c *cmdbController) DeleteHost(ctx *gin.Context) {
	instanceid := ctx.Param("instanceid")
	if err := v1.CoreV1.CMDB().Host().DeleteHost(ctx, instanceid); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func (c *cmdbController) DeleteHosts(ctx *gin.Context) {
	params := &dto.InstancesReq{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.CMDB().Host().DeleteHosts(ctx, params.Ids); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func (c *cmdbController) WebShell(ctx *gin.Context) {
	// 设置默认xterm窗口大小
	cols, _ := strconv.Atoi(ctx.DefaultQuery("cols", "188"))
	rows, _ := strconv.Atoi(ctx.DefaultQuery("rows", "42"))
	err := v1.CoreV1.CMDB().Host().WebShell(ctx.Writer, ctx.Request, cols, rows)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		return
	}
	v1.Log.Info("web shell connect success")
}
