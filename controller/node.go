package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var Node node

type node struct{}

func NodeRegister(router *gin.RouterGroup) {
	router.GET("/list", Node.GetNodeList)
	router.GET("/detail", Node.GetNodeDetail)
}

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
//@Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/node/list [get]
func (n *node) GetNodeList(ctx *gin.Context) {
	params := &dto.NodeListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.Node.GetNodes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取node列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
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
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/node/detail [get]
func (n *node) GetNodeDetail(ctx *gin.Context) {
	params := &dto.NodeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.Node.GetNodeDetail(params.Name)
	if err != nil {
		logger.Error("获取node详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
