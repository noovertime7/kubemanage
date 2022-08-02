package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
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
func (p *pod) UpdatePod(ctx *gin.Context) {
	params := &dto.PodUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, err)
		return
	}
	if err := service.Pod.UpdatePod(params.PodName, params.NameSpace, params.Content); err != nil {
		logger.Error("POD更新失败", err.Error())
		middleware.ResponseError(ctx, 10005, err)
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// GetPodContainer 获取Pod内容器名
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
func (p *pod) GetPodNumPreNp(ctx *gin.Context) {
	data, err := service.Pod.GetPodNumPerNp()
	if err != nil {
		middleware.ResponseError(ctx, 10006, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
