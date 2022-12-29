package operation

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

// GetOperationRecordList
// @Tags      SysOperationRecord
// @Summary   分页获取SysOperationRecord列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     dto.OperationListInput                        true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  middleware.Response{data=dto.OperationListOutPut,msg=string}  "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/operation/get_operations [get]
func (o *operationController) GetOperationRecordList(ctx *gin.Context) {
	params := &dto.OperationListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := v1.CoreV1.System().Operation().GetPageList(ctx, params)
	if err != nil {
		v1.Log.ErrorWithErr("查询失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// DeleteOperationRecord
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     dto.Empty
// @Success   200   {object}  middleware.Response{msg=string}  "删除SysOperationRecord"
// @Router    /api/operation/{id}/delete_operation [delete]
func (o *operationController) DeleteOperationRecord(ctx *gin.Context) {
	recordId, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.System().Operation().DeleteRecord(ctx, recordId); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// DeleteOperationRecords
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     dto.IdsReq
// @Success   200   {object}  middleware.Response{msg=string}  "删除SysOperationRecord"
// @Router    /api/operation/delete_operations [delete]
func (o *operationController) DeleteOperationRecords(ctx *gin.Context) {
	params := &dto.IdsReq{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().Operation().DeleteRecords(ctx, params.Ids); err != nil {
		v1.Log.ErrorWithErr("批量删除失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}
