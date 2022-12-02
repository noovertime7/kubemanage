package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/cmd/app/options"
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
		operation.NewOperationRouter(apiGroup)
		user.NewUserRouter(apiGroup)
	}
	installOperationRouters(apiGroup)
}

func installOperationRouters(apiGroup *gin.RouterGroup) {
	// 需要操作记录
	apiGroup.Use(middleware.OperationRecord())
	{
		other.NewSwaggarRoute(apiGroup)
		kubeController.NewKubeRouter(apiGroup)
		menu.NewMenuRouter(apiGroup)
		authority.NewCasbinRouter(apiGroup)
	}
}
