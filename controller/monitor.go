package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

type monitor struct {
	service *service.MonitorService
}

func MonitroRegister(router *gin.RouterGroup) {
	m := &monitor{service: service.NewMonitorService()}
	router.GET("/image_list", m.GetImageList)
}

func (m *monitor) GetImageList(ctx *gin.Context) {
	params := &dto.ImageListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := m.service.GetClusterImageList(params)
	if err != nil {
		logger.Error("获取镜像列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
