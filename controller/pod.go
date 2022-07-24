package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubemanage/service"
	"net/http"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  ":绑定参数失败:" + err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPods(parmas.FilterName, parmas.NameSpace, parmas.Limit, parmas.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取POD列表成功",
		"data": data,
	})
}
