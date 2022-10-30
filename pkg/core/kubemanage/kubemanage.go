package kubemanage

import (
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
)

var CoreV1 v1.CoreService

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	CoreV1 = v1.New(o.ComponentConfig, o.Factory)
}
