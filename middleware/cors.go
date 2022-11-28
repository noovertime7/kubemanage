package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cores 处理跨域请求，支持options访问
func Cores() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"},
		MaxAge:          6 * time.Hour,
	}

	return cors.New(c)
}
