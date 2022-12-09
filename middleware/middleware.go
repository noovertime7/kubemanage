package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
	"k8s.io/apimachinery/pkg/util/sets"
)

var AlwaysAllowPath sets.String

func InstallMiddlewares(ginEngine *gin.RouterGroup) {
	// 初始化可忽略的请求路径
	AlwaysAllowPath = sets.NewString(pkg.LoginURL, pkg.LogoutURL, pkg.WebShellURL)
	ginEngine.Use(Logger(), Cores(), Limiter(), Recovery(true), TranslationMiddleware(), JWTAuth(), CasbinHandler())
}
