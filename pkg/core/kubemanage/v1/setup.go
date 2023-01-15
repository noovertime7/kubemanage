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
	if config.SysConfig.CMDB.HostCheck.HostCheckEnable {
		startChecker()
	}
}

func startChecker() {
	// 启动checker factory
	CoreV1.CMDB().StartChecker()
	// 启动生产者
	handler := wait.NewDefaultBackoff(time.Duration(config.SysConfig.CMDB.HostCheck.HostCheckDuration) * time.Minute)
	Log.Infof("start host check every %d minutes...", config.SysConfig.CMDB.HostCheck.HostCheckDuration)
	go func() {
		wait.BackoffUntil(func() {
			if err := CoreV1.CMDB().Host().StartHostCheck(); err != nil {
				Log.ErrorWithErr("host check err", err)
				return
			}
		}, handler, true, runtime.SystemContext.Done())
	}()
}
