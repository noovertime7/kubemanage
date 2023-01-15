package v1

import (
	"time"

	"github.com/noovertime7/kubemanage/runtime"

	"github.com/noovertime7/kubemanage/runtime/wait"

	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/pkg/logger"
)

var CoreV1 CoreService

var Log Logger

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 5 seconds.
	quit := runtime.SetupSignalHandler()

	// Setup System Context
	runtime.SetupContext(quit)

	Log = logger.New(logger.LG)
	CoreV1 = New(config.SysConfig, o.Factory)
	startChecker()
}

func startChecker() {
	// 启动checker factory
	CoreV1.CMDB().StartChecker()
	// 启动生产者
	handler := wait.NewDefaultBackoff(60 * time.Second)
	go func() {
		wait.BackoffUntil(func() {
			CoreV1.CMDB().Host().StartHostCheck()
		}, handler, true, runtime.SystemContext.Done())
		return
	}()
}
