package kubeController

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var PersistentVolume persistentVolume

type persistentVolume struct{}

// DeletePersistentVolume 删除persistentVolume
// ListPage godoc
// @Summary      删除persistentVolume
// @Description  删除persistentVolume
// @Tags         PersistentVolume
// @ID           /api/k8s/spersistentvolume/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "persistentvolume名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/spersistentvolume/del [delete]
func (n *persistentVolume) DeletePersistentVolume(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.PersistentVolume.DeletePersistentVolume(params.Name); err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
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
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.PersistentVolumeResp}"
// @Router       /api/k8s/persistentvolume/list [get]
func (n *persistentVolume) GetPersistentVolumeList(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.PersistentVolume.GetPersistentVolumes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
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
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": *coreV1.PersistentVolume}"
// @Router       /api/k8s/persistentvolume/detail [get]
func (n *persistentVolume) GetPersistentVolumeDetail(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.PersistentVolume.GetPersistentVolumesDetail(params.Name)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
