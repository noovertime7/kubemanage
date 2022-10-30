package main

import (
	"github.com/noovertime7/kubemanage/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
