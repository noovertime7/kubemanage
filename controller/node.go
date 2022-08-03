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
