package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/wonderivan/logger"
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
	params := &kubernetes.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.DaemonSet.DeleteDaemonSet(params.Name, params.NameSpace); err != nil {
		logger.Error("删除DaemonSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
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
	params := &kubernetes.DaemonSetUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.DaemonSet.UpdateDaemonSet(params.Content, params.NameSpace); err != nil {
		logger.Error("更新DaemonSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
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
	params := &kubernetes.DaemonSetListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.DaemonSet.GetDaemonSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取DaemonSet列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
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
	params := &kubernetes.DaemonSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.DaemonSet.GetDaemonSetDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取DaemonSet详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
