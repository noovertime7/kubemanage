package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
	_ "k8s.io/api/core/v1"
)

var Pod pod

type pod struct{}

func PodRegister(router *gin.RouterGroup) {
	router.GET("/pods", Pod.GetPods)
	router.GET("/pod/detail", Pod.GetPodDetail)
	router.DELETE("/pod/del", Pod.DeletePod)
	router.PUT("/pod/update", Pod.UpdatePod)
	router.GET("/pod/container", Pod.GetPodContainer)
	router.GET("/pod/log", Pod.GetPodLog)
	router.GET("/pod/numnp", Pod.GetPodNumPreNp)

}

// GetPods 获取pod，支持分页过滤排序
// GetPods 获取pod，支持分页过滤排序
// ListPage godoc
// @Summary      获取pod列表
// @Description  获取pod列表
// @Tags         pod
// @ID           /api/k8s/pods
// @Accept       json
// @Produce      json
// @Param        pod_name   query     string                                      true  "POD名称"
// @Param        namespace  query     string                                      true  "命名空间"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data": service.PodsResp}"
// @Router       /api/k8s/pods [get]
func (p *pod) GetPods(ctx *gin.Context) {
	//处理入参
	parmas := &dto.PodListInput{}
	if err := parmas.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10001, err)
		return
	}
	data, err := service.Pod.GetPods(parmas.FilterName, parmas.NameSpace, parmas.Limit, parmas.Page)
	if err != nil {
		middleware.ResponseError(ctx, 10002, err)
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
// @Param        pod_name   query     string                        true  "POD名称"
// @Param        namespace  query     string                        true  "命名空间"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Pod }"
// @Router       /api/k8s/pod/detail [get]
func (p *pod) GetPodDetail(ctx *gin.Context) {
	//处理入参
	parmas := &dto.PodNameNsInput{}
	if err := parmas.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	data, err := service.Pod.GetPodDetail(parmas.PodName, parmas.NameSpace)
	if err != nil {
		middleware.ResponseError(ctx, 10004, err)
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
// @Param        pod_name   query     string                        true  "POD名称"
// @Param        namespace  query     string                        true  "命名空间"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/del [delete]
func (p *pod) DeletePod(ctx *gin.Context) {
	params := &dto.PodNameNsInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	if err := service.Pod.DeletePod(params.PodName, params.NameSpace); err != nil {
		middleware.ResponseError(ctx, 10004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// UpdatePod 更新POD
// ListPage godoc
// @Summary      更新POD
// @Description  更新POD
// @Tags         pod
// @ID           /api/k8s/pod/update
// @Accept       json
// @Produce      json
// @Param        pod_name   query     string                        true  "POD名称"
// @Param        namespace  query     string                        true  "命名空间"
// @Param        content    query     string                        true  "更新内容"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/update [put]
func (p *pod) UpdatePod(ctx *gin.Context) {
	params := &dto.PodUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	if err := service.Pod.UpdatePod(params.NameSpace, params.Content); err != nil {
		logger.Error("POD更新失败", err.Error())
		middleware.ResponseError(ctx, 10005, err)
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// GetPodContainer 获取Pod内容器名
// ListPage godoc
// @Summary      获取Pod内容器名
// @Description  获取Pod内容器名
// @Tags         pod
// @ID           /api/k8s/pod/container
// @Accept       json
// @Produce      json
// @Param        pod_name   query     string                              true  "POD名称"
// @Param        namespace  query     string                              true  "命名空间"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/container [get]
func (p *pod) GetPodContainer(ctx *gin.Context) {
	params := &dto.PodNameNsInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	data, err := service.Pod.GetPodContainer(params.PodName, params.NameSpace)
	if err != nil {
		middleware.ResponseError(ctx, 10004, err)
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
// @Param        pod_name        query     string                            true  "POD名称"
// @Param        namespace       query     string                            true  "命名空间"
// @Param        container_name  query     string                            true  "容器名"
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/pod/log [get]
func (p *pod) GetPodLog(ctx *gin.Context) {
	params := &dto.PodGetLogInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	data, err := service.Pod.GetPodLog(params.ContainerName, params.PodName, params.NameSpace)
	if err != nil {
		logger.Error("POD更新失败", err.Error())
		middleware.ResponseError(ctx, 10005, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

//GetPodNumPreNp 根据命名空间获取数量
// ListPage godoc
// @Summary      根据命名空间获取数量
// @Description  根据命名空间获取数量
// @Tags         pod
// @ID           /api/k8s/pod/numnp
// @Accept       json
// @Produce      json
//@Success      200        {object}  middleware.Response"{"code": 200, msg="","data":[]service.PodsNp }"
// @Router       /api/k8s/pod/numnp [get]
func (p *pod) GetPodNumPreNp(ctx *gin.Context) {
	data, err := service.Pod.GetPodNumPerNp()
	if err != nil {
		middleware.ResponseError(ctx, 10006, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
