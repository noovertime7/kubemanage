package middleware

import "github.com/gin-gonic/gin"

func InstallMiddlewares(ginEngine *gin.RouterGroup) {
	ginEngine.Use(Cores(), Limiter(), TranslationMiddleware(), JWTAuth(), CasbinHandler())
}
