package main

import (
	"github.com/noovertime7/kubemanage/router"
	"github.com/noovertime7/kubemanage/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//初始化K8s client
	service.K8s.Init()
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
