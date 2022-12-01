package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/wonderivan/logger"
)

var Node node

type node struct{}

// GetNodeList 获取Node列表
// ListPage godoc
// @Summary      获取Node列表
// @Description  获取Node列表
// @Tags         Node
// @ID           /api/k8s/node/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/node/list [get]
func (n *node) GetNodeList(ctx *gin.Context) {
	params := &kubernetes.NodeListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Node.GetNodes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取node列表失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetNodeDetail 获取Node详情
// ListPage godoc
// @Summary      获取Node详情
// @Description  获取Node详情
// @Tags         Node
// @ID           /api/k8s/node/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "node名称"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/node/detail [get]
func (n *node) GetNodeDetail(ctx *gin.Context) {
	params := &kubernetes.NodeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Node.GetNodeDetail(params.Name)
	if err != nil {
		logger.Error("获取node详情失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
