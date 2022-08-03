package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var DaemonSet daemonSet

type daemonSet struct{}

func DaemonSetRegister(router *gin.RouterGroup) {
	router.DELETE("/del", DaemonSet.DeleteDaemonSet)
	router.PUT("/update", DaemonSet.UpdateDaemonSet)
	router.GET("/list", DaemonSet.GetDaemonSetList)
	router.GET("/detail", DaemonSet.GetDaemonSetDetail)
}

func (s *daemonSet) DeleteDaemonSet(ctx *gin.Context) {
	params := &dto.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.DaemonSet.DeleteDaemonSet(params.Name, params.NameSpace); err != nil {
		logger.Error("删除DaemonSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (s *daemonSet) UpdateDaemonSet(ctx *gin.Context) {
	params := &dto.DaemonSetUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.DaemonSet.UpdateDaemonSet(params.Content, params.NameSpace); err != nil {
		logger.Error("更新DaemonSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

func (s *daemonSet) GetDaemonSetList(ctx *gin.Context) {
	params := &dto.DaemonSetListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.DaemonSet.GetDaemonSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取DaemonSet列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *daemonSet) GetDaemonSetDetail(ctx *gin.Context) {
	params := &dto.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.DaemonSet.GetDaemonSetDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取DaemonSet详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
