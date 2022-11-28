package v1

import (
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/cmd/app/options"
)

var CoreV1 CoreService

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	CoreV1 = New(config.SysConfig, o.Factory)
}
