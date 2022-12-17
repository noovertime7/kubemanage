package api

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

// GetApiList
// @Tags      API
// @Summary   获取API列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     dto.Empty
// @Success   200   {object}  middleware.Response{data=[]model.SysApi,msg=string}  "获取API列表"
// @Router    /api/sysApi/getAPiList [get]
func (a *apiController) GetApiList(ctx *gin.Context) {
	data, err := v1.CoreV1.System().Api().GetApiList(ctx)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
	}
	middleware.ResponseSuccess(ctx, data)
}
