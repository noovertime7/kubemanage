package authority

import (
	"github.com/gin-gonic/gin"
)

type casbinController struct{}

func NewCasbinRouter(ginEngine *gin.RouterGroup) {
	cas := casbinController{}
	cas.initRoutes(ginEngine)
}

func (c *casbinController) initRoutes(ginEngine *gin.RouterGroup) {
	casRoute := ginEngine.Group("/casbin")
	cas := &casbinController{}
	{
		casRoute.GET("/getPolicyPathByAuthorityId", cas.GetPolicyPathByAuthorityId)
	}
}
