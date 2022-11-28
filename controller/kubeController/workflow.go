package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/wonderivan/logger"
)

var WorkFlow workflow

type workflow struct{}

// CreateWorkFlow 创建workflow
// ListPage godoc
// @Summary      创建workflow
// @Description  创建workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.WorkFlowCreateInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/workflow/create [post]
func (w *workflow) CreateWorkFlow(ctx *gin.Context) {
	params := &kubernetes.WorkFlowCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 60001, err)
		return
	}
	if err := v1.CoreV1.WorkFlow().Save(ctx, params); err != nil {
		logger.Error("创建失败", err)
		middleware.ResponseError(ctx, 60002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// DeleteWorkflow 删除Workflow
// ListPage godoc
// @Summary      删除Workflow
// @Description  删除Workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/del
// @Accept       json
// @Produce      json
// @Param        ID       query  int  true  "Workflow ID"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/workflow/del [delete]
func (w *workflow) DeleteWorkflow(ctx *gin.Context) {
	params := &kubernetes.WorkFlowIDInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 60001, err)
		return
	}
	if err := v1.CoreV1.WorkFlow().Delete(ctx, params.ID); err != nil {
		logger.Error("删除失败", err)
		middleware.ResponseError(ctx, 60002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// GetWorkflowList 查看Configmap列表
// ListPage godoc
// @Summary      查看Configmap列表
// @Description  查看Configmap列表
// @Tags         Workflow
// @ID           /api/k8s/workflow/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/workflow/list [get]
func (w *workflow) GetWorkflowList(ctx *gin.Context) {
	params := &kubernetes.WorkFlowListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 60001, err)
		return
	}
	data, err := v1.CoreV1.WorkFlow().FindList(ctx, params)
	if err != nil {
		logger.Error("查询失败", err)
		middleware.ResponseError(ctx, 60002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetWorkflowByID 根据ID查看workflow
// ListPage godoc
// @Summary      根据ID查看workflow
// @Description  根据ID查看workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/id
// @Accept       json
// @Produce      json
// @Param        ID       query  int  true  "Workflow ID"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/workflow/id [get]
func (w *workflow) GetWorkflowByID(ctx *gin.Context) {
	params := &kubernetes.WorkFlowIDInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 60001, err)
		return
	}
	data, err := v1.CoreV1.WorkFlow().Find(ctx, params)
	if err != nil {
		logger.Error("查询失败", err)
		middleware.ResponseError(ctx, 60002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
