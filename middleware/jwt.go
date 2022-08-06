package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
	"github.com/pkg/errors"
)

// JWTAuth jwt认证函数
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Request.URL.String()) == 15 && context.Request.URL.String()[0:15] == "/api/user/login" {
			context.Next()
			return
		}
		// 处理验证逻辑
		token := context.Request.Header.Get("token")
		if token == "" {
			ResponseError(context, 11000, errors.New("请求未携带token,无权限访问"))
			context.Abort()
			return
		}
		// 解析token内容
		claims, err := public.JWTToken.ParseToken(token)
		if err != nil {
			// token过期错误
			if err.Error() == "TokenExpired" {
				ResponseError(context, 11001, errors.New("token过期"))
				context.Abort()
				return
			}
			// 解析其他错误
			ResponseError(context, 11002, errors.New("token解析失败"))
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
		context.Next()
	}
}
