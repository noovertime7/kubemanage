package main

import (
	"github.com/gin-gonic/gin"
	"kubemanage/config"
	"kubemanage/controller"
	"kubemanage/service"
)

func main() {
	//初始化K8s client
	service.K8s.Init()
	//初始化gin
	r := gin.Default()
	controller.Router.InitApiRouter(r)
	//启动server
	r.Run(config.ListenAddr)
}
