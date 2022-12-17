package router

import (
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/controller/api"
	"github.com/noovertime7/kubemanage/controller/authority"
	"github.com/noovertime7/kubemanage/controller/kubeController"
	"github.com/noovertime7/kubemanage/controller/menu"
	"github.com/noovertime7/kubemanage/controller/operation"
	"github.com/noovertime7/kubemanage/controller/other"
	"github.com/noovertime7/kubemanage/controller/user"
	"github.com/noovertime7/kubemanage/middleware"
)

func InstallRouters(opt *options.Options) {
	apiGroup := opt.GinEngine.Group("/api")
	middleware.InstallMiddlewares(apiGroup)
	//安装不需要操作记录路由
	{
		api.NewApiRouter(apiGroup)
		operation.NewOperationRouter(apiGroup)
		user.NewUserRouter(apiGroup)
		other.NewSwaggarRoute(apiGroup)
		kubeController.NewKubeRouter(apiGroup)
		menu.NewMenuRouter(apiGroup)
		authority.NewCasbinRouter(apiGroup)
	}
}
