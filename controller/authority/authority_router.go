package authority

import (
	"github.com/gin-gonic/gin"
)

type authorityController struct{}

func NewCasbinRouter(ginEngine *gin.RouterGroup) {
	cas := authorityController{}
	cas.initRoutes(ginEngine)
}

func (a *authorityController) initRoutes(ginEngine *gin.RouterGroup) {
	casRoute := ginEngine.Group("/authority")
	cas := &authorityController{}
	{
		casRoute.GET("/getPolicyPathByAuthorityId", cas.GetPolicyPathByAuthorityId)
		casRoute.GET("/getAuthorityList", cas.GetAuthorityList)
	}
}
