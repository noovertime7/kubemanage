package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var PersistentVolume persistentVolume

type persistentVolume struct{}

func PersistentVolumeRegister(router *gin.RouterGroup) {
	router.DELETE("/del", PersistentVolume.DeletePersistentVolume)
	router.GET("/list", PersistentVolume.GetPersistentVolumeList)
	router.GET("/detail", PersistentVolume.GetPersistentVolumeDetail)
}

// DeletePersistentVolume 删除persistentVolume
// ListPage godoc
// @Summary      删除persistentVolume
// @Description  删除persistentVolume
// @Tags         PersistentVolume
// @ID           /api/k8s/spersistentvolume/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "persistentvolume名称"
//@Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/spersistentvolume/del [delete]
func (n *persistentVolume) DeletePersistentVolume(ctx *gin.Context) {
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

// GetPersistentVolumeList 获取PV列表
// ListPage godoc
// @Summary      获取PV列表
// @Description  获取PV列表
// @Tags         PersistentVolume
// @ID           /api/k8s/persistentvolume/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  true  "persistentVolume名称"
// @Param        page         query  int     true  "页码"
// @Param        limit        query  int     true  "分页限制"
//@Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.PersistentVolumeResp}"
// @Router       /api/k8s/persistentvolume/list [get]
func (n *persistentVolume) GetPersistentVolumeList(ctx *gin.Context) {
	params := &dto.PersistentVolumeListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取persistentVolume列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetPersistentVolumeDetail 获取PV的详细信息
// ListPage godoc
// @Summary      获取PV的详细信息
// @Description  获取PV的详细信息
// @Tags         PersistentVolume
// @ID           /api/k8s/persistentvolume/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "persistentVolume名称"
//@Success       200  {object}  middleware.Response"{"code": 200, msg="","data": *coreV1.PersistentVolume}"
// @Router       /api/k8s/persistentvolume/detail [get]
func (n *persistentVolume) GetPersistentVolumeDetail(ctx *gin.Context) {
	params := &dto.PersistentVolumeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumesDetail(params.Name)
	if err != nil {
		logger.Error("获取persistentVolume详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
