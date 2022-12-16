package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if AlwaysAllowPath.Has(c.Request.URL.Path) {
			return
		}
		waitUse, err := utils.GetClaims(c)
		if err != nil {
			ResponseError(c, globalError.NewGlobalError(globalError.ServerError, err))
			c.Abort()
			return
		}
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := v1.CoreV1.System().CasbinService().Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			ResponseError(c, globalError.NewGlobalError(globalError.AuthErr, fmt.Errorf("角色ID %d 请求 %s %s 无权限", waitUse.AuthorityId, act, obj)))
			c.Abort()
			return
		}
	}
}
