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
	}
	{
		cmdbRoute.POST("/deleteHosts", c.DeleteHosts)
		cmdbRoute.DELETE("/:instanceid/deleteHost", c.DeleteHost)
		cmdbRoute.POST("/createHost", c.CreateHost)
		cmdbRoute.POST("/updateHost", c.UpdateHost)
	}
}
