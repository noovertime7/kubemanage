package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/cmd/app"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
