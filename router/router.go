package router

import (
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/controller/authority"
	"github.com/noovertime7/kubemanage/controller/kubeController"
	"github.com/noovertime7/kubemanage/controller/menu"
	"github.com/noovertime7/kubemanage/controller/other"
	"github.com/noovertime7/kubemanage/controller/user"
	"github.com/noovertime7/kubemanage/middleware"
)

func InstallRouters(opt *options.Options) {
	apiGroup := opt.GinEngine.Group("/api")

	// install middlewares
	middleware.InstallMiddlewares(apiGroup)

	other.NewSwaggarRoute(apiGroup)
	user.NewUserRouter(apiGroup)
	kubeController.NewKubeRouter(apiGroup)
	menu.NewMenuRouter(apiGroup)
	authority.NewCasbinRouter(apiGroup)
}
