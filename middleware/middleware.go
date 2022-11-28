package middleware

import "github.com/gin-gonic/gin"

func InstallMiddlewares(ginEngine *gin.RouterGroup) {
	ginEngine.Use(Logger(), Cores(), Limiter(), Recovery(true), TranslationMiddleware(), JWTAuth(), CasbinHandler())
}
