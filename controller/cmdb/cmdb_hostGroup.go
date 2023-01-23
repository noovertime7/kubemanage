package cmdb

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

func (c *cmdbController) CreateHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdate{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().CreateHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

func (c *cmdbController) CreateSonHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdate{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().CreateSonHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

func (c *cmdbController) UpdateHostGroup(ctx *gin.Context) {
	params := &dto.HostGroupCreateOrUpdate{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	err := v1.CoreV1.CMDB().HostGroup().UpdateHostGroup(ctx, params)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

func (c *cmdbController) DeleteHostGroup(ctx *gin.Context) {
	instanceID := ctx.Param("instanceID")
	if err := v1.CoreV1.CMDB().HostGroup().DeleteHostGroup(ctx, instanceID); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func (c *cmdbController) GetHostGroupTree(ctx *gin.Context) {
	data, err := v1.CoreV1.CMDB().HostGroup().GetHostGroupTree(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (c *cmdbController) GetHostGroupList(ctx *gin.Context) {
	data, err := v1.CoreV1.CMDB().HostGroup().GetHostGroupList(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
