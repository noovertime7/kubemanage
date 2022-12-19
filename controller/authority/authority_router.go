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
	{
		casRoute.GET("/getPolicyPathByAuthorityId", a.GetPolicyPathByAuthorityId)
		casRoute.POST("/updateCasbinByAuthority", a.UpdateCasbinByAuthorityId)
		casRoute.GET("/getAuthorityList", a.GetAuthorityList)
		casRoute.POST("/createAuthority", a.CreateAuthority)
		casRoute.DELETE("/:authID/delAuthority", a.DeleteAuthority)
		casRoute.PUT("/updateAuthority", a.UpdateAuthority)
	}
}
