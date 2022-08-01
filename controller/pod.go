package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var Pod pod

type pod struct{}

func PodRegister(router *gin.RouterGroup) {
	p := pod{}
	router.GET("k8s/pods", p.GetPods)
}

// GetPods 获取pod，支持分页过滤排序
func (p *pod) GetPods(ctx *gin.Context) {
	//处理入参
	parmas := &dto.PodListInput{}
	if err := parmas.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10001, errors.New("绑定参数失败"))
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
	parmas := new(struct {
		PodName   string `form:"pod_name"`
		NameSpace string `form:"namespace"`
	})
	if err := ctx.Bind(parmas); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10003, errors.New("绑定参数失败"))
		return
	}
	data, err := service.Pod.GetPodDetail(parmas.PodName, parmas.NameSpace)
	if err != nil {
		middleware.ResponseError(ctx, 10004, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
