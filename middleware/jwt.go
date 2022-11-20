package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

// JWTAuth jwt认证函数
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Request.URL.String()) == 15 && context.Request.URL.String()[0:15] == "/api/user/login" {
			context.Next()
			return
		}
		// 处理验证逻辑
		claims, err := utils.GetClaims(context)
		if err != nil {
			ResponseError(context, 405, err)
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
		context.Next()
	}
}
