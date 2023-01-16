package cmdb

import (
	"github.com/gin-gonic/gin"
)

type cmdbController struct{}

func NewCMDBRouter(ginEngine *gin.RouterGroup) {
	cmdb := cmdbController{}
	cmdb.initRoutes(ginEngine)
}

func (c *cmdbController) initRoutes(ginEngine *gin.RouterGroup) {
	cmdbRoute := ginEngine.Group("/cmdb")
	{
		cmdbRoute.GET("/getHostGroupTree", c.GetHostGroupTree)
		cmdbRoute.GET("/getHostGroupList", c.GetHostGroupList)
	}
	{
		cmdbRoute.GET("/:groupID/pageHost", c.PageHost)
		cmdbRoute.POST("/deleteHosts", c.DeleteHosts)
		cmdbRoute.DELETE("/:instanceid/deleteHost", c.DeleteHost)
		cmdbRoute.POST("/createHost", c.CreateHost)
		cmdbRoute.POST("/updateHost", c.UpdateHost)
	}
	{
		cmdbRoute.POST("/createSecret", c.CreateSecret)
		cmdbRoute.POST("/updateSecret", c.UpdateSecret)
		cmdbRoute.GET("/pageSecret", c.PageSecret)
		cmdbRoute.POST("/deleteSecrets", c.DeleteSecrets)
		cmdbRoute.DELETE("/:instanceid/deleteSecret", c.DeleteSecret)
	}
}
