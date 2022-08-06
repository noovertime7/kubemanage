package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cores 处理跨域请求，支持options访问
func Cores() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取请求方法
		method := context.Request.Method

		// 添加跨域响应头
		context.Header("Content-Type", "application/json")
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Max-Age", "86400")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		context.Header("Access-Control-Allow-Credentials", "false")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		context.Next()
	}

}
