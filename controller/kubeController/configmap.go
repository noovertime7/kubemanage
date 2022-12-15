package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"

	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var Configmap configmap

type configmap struct{}

// DeleteConfigmap 删除Configmap
// ListPage godoc
// @Summary      删除Configmap
// @Description  删除Configmap
// @Tags         Configmap
// @ID           /api/k8s/configmap/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Configmap名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/configmap/del [delete]
func (s *configmap) DeleteConfigmap(ctx *gin.Context) {
	params := &kubeDto.ConfigmapNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Configmap.DeleteConfigmap(params.Name, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateConfigmap 更新Configmap
// ListPage godoc
// @Summary      更新Configmap
// @Description  更新Configmap
// @Tags         Configmap
// @ID           /api/k8s/configmap/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/configmap/update [put]
func (s *configmap) UpdateConfigmap(ctx *gin.Context) {
	params := &kubeDto.ConfigmapUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Configmap.UpdateConfigmap(params.Content, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetConfigmapList 查看Configmap列表
// ListPage godoc
// @Summary      查看Configmap列表
// @Description  查看Configmap列表
// @Tags         Configmap
// @ID           /api/k8s/configmap/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/configmap/list [get]
func (s *configmap) GetConfigmapList(ctx *gin.Context) {
	params := &kubeDto.ConfigmapListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Configmap.GetConfigmaps(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetConfigmapDetail 获取Configmap详情
// ListPage godoc
// @Summary      获取Configmap详情
// @Description  获取Configmap详情
// @Tags         Configmap
// @ID           /api/k8s/configmap/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Configmap名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/configmap/detail [get]
func (s *configmap) GetConfigmapDetail(ctx *gin.Context) {
	params := &kubeDto.ConfigmapNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Configmap.GetConfigmapDetail(params.Name, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
