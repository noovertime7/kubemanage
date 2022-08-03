package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var PersistentVolume persistentvolume

type persistentvolume struct{}

func PersistentVolumeRegister(router *gin.RouterGroup) {
	router.DELETE("/del", PersistentVolume.DeletePersistentVolume)
	router.GET("/list", PersistentVolume.GetPersistentVolumeList)
	router.GET("/detail", PersistentVolume.GetPersistentVolumeDetail)
}

func (n *persistentvolume) DeletePersistentVolume(ctx *gin.Context) {
	params := &dto.PersistentVolumeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.PersistentVolume.DeletePersistentVolume(params.Name); err != nil {
		logger.Error("删除persistentvolume失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (n *persistentvolume) GetPersistentVolumeList(ctx *gin.Context) {
	params := &dto.PersistentVolumeListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取persistentvolume列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (n *persistentvolume) GetPersistentVolumeDetail(ctx *gin.Context) {
	params := &dto.PersistentVolumeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumesDetail(params.Name)
	if err != nil {
		logger.Error("获取persistentvolume详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
