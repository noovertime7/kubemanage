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
		cmdbRoute.POST("/createHostGroup", c.CreateHostGroup)
		cmdbRoute.POST("/createSonHostGroup", c.CreateSonHostGroup)
		cmdbRoute.PUT("/updateHostGroup", c.UpdateHostGroup)
		cmdbRoute.DELETE("/:instanceID/deleteHostGroup", c.DeleteHostGroup)
		cmdbRoute.GET("/getHostGroupTree", c.GetHostGroupTree)
		cmdbRoute.GET("/getHostGroupList", c.GetHostGroupList)
	}
	{
		cmdbRoute.GET("/:groupID/pageHost", c.PageHost)
		cmdbRoute.GET("/getHostsList", c.GetHostList)
		cmdbRoute.POST("/deleteHosts", c.DeleteHosts)
		cmdbRoute.DELETE("/:instanceID/deleteHost", c.DeleteHost)
		cmdbRoute.POST("/createHost", c.CreateHost)
		cmdbRoute.POST("/updateHost", c.UpdateHost)
		cmdbRoute.GET("/webshell", c.WebShell)
	}
	{
		cmdbRoute.POST("/createSecret", c.CreateSecret)
		cmdbRoute.POST("/updateSecret", c.UpdateSecret)
		cmdbRoute.GET("/pageSecret", c.PageSecret)
		cmdbRoute.GET("/getSecretList", c.GetSecretList)
		cmdbRoute.POST("/deleteSecrets", c.DeleteSecrets)
		cmdbRoute.DELETE("/:instanceID/deleteSecret", c.DeleteSecret)
	}
}
