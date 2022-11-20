package middleware

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/pkg/errors"
	"strconv"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.URL.String()) == 15 && c.Request.URL.String()[0:15] == "/api/user/login" {
			c.Next()
			return
		}
		waitUse, err := utils.GetClaims(c)
		if err != nil {
			ResponseError(c, 405, err)
			c.Abort()
			return
		}
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := v1.CoreV1.CasbinService().Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			ResponseError(c, 403, errors.New("403 forbidden"))
			c.Abort()
			return
		}
	}
}
