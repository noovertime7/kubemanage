package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubemanage/middleware"
	"kubemanage/service"
)

var Pod pod

type pod struct{}

// GetPods 获取pod，支持分页过滤排序
func (p *pod) GetPods(ctx *gin.Context) {
	//处理入参
	parmas := new(struct {
		FilterName string `form:"filter_name"`
		NameSpace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	if err := ctx.Bind(parmas); err != nil {
		logger.Error("绑定参数失败:", err.Error())
		middleware.ResponseError(ctx, 10001, errors.New("绑定参数失败"))
		return
	}
	data, err := service.Pod.GetPods(parmas.FilterName, parmas.NameSpace, parmas.Limit, parmas.Page)
	if err != nil {
		middleware.ResponseError(ctx, 10002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "获取POD成功", data)
}
