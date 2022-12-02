package v1

import (
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/pkg/logger"
)

var CoreV1 CoreService

var Log Logger

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	Log = logger.NewErrorLoger()
	CoreV1 = New(config.SysConfig, o.Factory)
}
