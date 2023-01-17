package middleware

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/noovertime7/kubemanage/pkg"
)

var AlwaysAllowPath sets.String

func InstallMiddlewares(ginEngine *gin.RouterGroup) {
	// 初始化可忽略的请求路径
	AlwaysAllowPath = sets.NewString(pkg.LoginURL, pkg.LogoutURL, pkg.WebShellURL, pkg.HostWebShell)
	ginEngine.Use(Logger(), Cores(), Limiter(), OperationRecord(), Recovery(true), TranslationMiddleware(), JWTAuth(), CasbinHandler())
}
