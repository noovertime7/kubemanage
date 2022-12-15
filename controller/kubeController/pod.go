package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	_ "k8s.io/api/core/v1"
)

var Pod pod

type pod struct{}

// GetPods 获取pod，支持分页过滤排序
// ListPage godoc
// @Summary      获取pod列表
// @Description  获取pod列表
// @Tags         pod
// @ID           /api/k8s/pods
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace    query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success      200 {object}  middleware.Response"{"code": 200, msg="","data": service.PodsResp}"
// @Router       /api/k8s/pods [get]
func (p *pod) GetPods(ctx *gin.Context) {
	//处理入参
	parmas := &kubeDto.PodListInput{}
	if err := parmas.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Pod.GetPods(parmas.FilterName, parmas.NameSpace, parmas.Limit, parmas.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetPodDetail 获取Pod详情
// ListPage godoc
// @Summary      获取Pod详情
// @Description  获取Pod详情
// @Tags         pod
// @ID           /api/k8s/pod/detail
// @Accept       json
// @Produce      json
// @Param        pod_name   query  string  true  "POD名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Pod }"
// @Router       /api/k8s/pod/detail [get]
func (p *pod) GetPodDetail(ctx *gin.Context) {
	//处理入参
	parmas := &kubeDto.PodNameNsInput{}
	if err := parmas.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Pod.GetPodDetail(parmas.PodName, parmas.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// DeletePod  删除POD
// ListPage godoc
// @Summary      删除POD
// @Description  删除POD
// @Tags         pod
// @ID           /api/k8s/pod/del
// @Accept       json
// @Produce      json
// @Param        pod_name   query  string  true  "POD名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/del [delete]
func (p *pod) DeletePod(ctx *gin.Context) {
	params := &kubeDto.PodNameNsInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Pod.DeletePod(params.PodName, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdatePod 更新POD
// ListPage godoc
// @Summary      更新POD
// @Description  更新POD
// @Tags         pod
// @ID           /api/k8s/pod/update
// @Accept       json
// @Produce      json
// @Param        pod_name   query  string  true  "POD名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/update [put]
func (p *pod) UpdatePod(ctx *gin.Context) {
	params := &kubeDto.PodUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Pod.UpdatePod(params.NameSpace, params.Content); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetPodContainer 获取Pod内容器名
// ListPage godoc
// @Summary      获取Pod内容器名
// @Description  获取Pod内容器名
// @Tags         pod
// @ID           /api/k8s/pod/container
// @Accept       json
// @Produce      json
// @Param        pod_name   query  string  true  "POD名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/container [get]
func (p *pod) GetPodContainer(ctx *gin.Context) {
	params := &kubeDto.PodNameNsInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Pod.GetPodContainer(params.PodName, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetPodLog 获取容器日志
// ListPage godoc
// @Summary      获取容器日志
// @Description  获取容器日志
// @Tags         pod
// @ID           /api/k8s/pod/log
// @Accept       json
// @Produce      json
// @Param        pod_name        query  string  true  "POD名称"
// @Param        namespace       query  string  true  "命名空间"
// @Param        container_name  query  string  true  "容器名"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/log [get]
func (p *pod) GetPodLog(ctx *gin.Context) {
	params := &kubeDto.PodGetLogInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Pod.GetPodLog(params.ContainerName, params.PodName, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetPodNumPreNp 根据命名空间获取数量
// ListPage godoc
// @Summary      根据命名空间获取数量
// @Description  根据命名空间获取数量
// @Tags         pod
// @ID           /api/k8s/pod/numnp
// @Accept       json
// @Produce      json
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":[]service.PodsNp }"
// @Router       /api/k8s/pod/numnp [get]
func (p *pod) GetPodNumPreNp(ctx *gin.Context) {
	data, err := kube.Pod.GetPodNumPerNp()
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (p *pod) WebShell(ctx *gin.Context) {
	ops := &kubeDto.WebShellOptions{}
	if err := ops.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
	}
	if err := v1.CoreV1.Cloud().Pods("").WebShellHandler(ops, ctx.Writer, ctx.Request); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
	}
}
