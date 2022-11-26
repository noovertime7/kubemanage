package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
	"github.com/pkg/errors"
)

func GetClaims(c *gin.Context) (*pkg.CustomClaims, error) {
	token := c.Request.Header.Get("token")
	if token == "" {
		return nil, errors.New("请求未携带token,无权限访问")
	}
	// 解析token内容
	claims, err := pkg.JWTToken.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) (uint, error) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0, err
		} else {
			return cl.AuthorityId, nil
		}
	} else {
		waitUse := claims.(*pkg.CustomClaims)
		return waitUse.AuthorityId, nil
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *pkg.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*pkg.CustomClaims)
		return waitUse
	}
}
