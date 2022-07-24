package controller

import (
	"github.com/gin-gonic/gin"
)

var Router router

type router struct{}

// InitApiRouter 初始化路由规则
func (r *router) InitApiRouter(router *gin.Engine) {
	router.GET("/api/k8s/pods", Pod.GetPods)
}
