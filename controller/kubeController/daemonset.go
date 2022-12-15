package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var DaemonSet daemonSet

type daemonSet struct{}

// DeleteDaemonSet 删除DaemonSet
// ListPage godoc
// @Summary      删除DaemonSet
// @Description  删除DaemonSet
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "DaemonSet名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/daemonset/del [delete]
func (s *daemonSet) DeleteDaemonSet(ctx *gin.Context) {
	params := &kubeDto.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.DaemonSet.DeleteDaemonSet(params.Name, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateDaemonSet 更新DaemonSet
// ListPage godoc
// @Summary      更新DaemonSet
// @Description  更新DaemonSet
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/daemonset/update [put]
func (s *daemonSet) UpdateDaemonSet(ctx *gin.Context) {
	params := &kubeDto.DaemonSetUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.DaemonSet.UpdateDaemonSet(params.Content, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetDaemonSetList 查看DaemonSet列表
// ListPage godoc
// @Summary      查看DaemonSet列表
// @Description  查看DaemonSet列表
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/daemonset/list [get]
func (s *daemonSet) GetDaemonSetList(ctx *gin.Context) {
	params := &kubeDto.DaemonSetListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.DaemonSet.GetDaemonSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetDaemonSetDetail 获取DaemonSet详情
// ListPage godoc
// @Summary      获取DaemonSet详情
// @Description  获取DaemonSet详情
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "DaemonSet名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/daemonset/detail [get]
func (s *daemonSet) GetDaemonSetDetail(ctx *gin.Context) {
	params := &kubeDto.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.DaemonSet.GetDaemonSetDetail(params.Name, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
