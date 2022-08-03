package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var NameSpace namespace

type namespace struct{}

func NameSpaceRegister(router *gin.RouterGroup) {
	router.DELETE("/del", NameSpace.DeleteNameSpace)
	router.GET("/list", NameSpace.GetNameSpaceList)
	router.GET("/detail", NameSpace.GetNameSpaceDetail)
}

func (n *namespace) DeleteNameSpace(ctx *gin.Context) {
	params := &dto.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.NameSpace.DeleteNameSpace(params.Name); err != nil {
		logger.Error("删除NameSpace失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (n *namespace) GetNameSpaceList(ctx *gin.Context) {
	params := &dto.NameSpaceListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.NameSpace.GetNameSpaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取Namespace列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (n *namespace) GetNameSpaceDetail(ctx *gin.Context) {
	params := &dto.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.NameSpace.GetNameSpacesDetail(params.Name)
	if err != nil {
		logger.Error("获取Namespace详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
