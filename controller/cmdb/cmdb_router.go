package cmdb

import (
	"github.com/gin-gonic/gin"
	"time"

	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
)

type cmdbController struct{}

func NewCMDBRouter(ginEngine *gin.RouterGroup) {
	cmdb := cmdbController{}
	cmdb.startChecker()
	cmdb.initRoutes(ginEngine)
}

func (c *cmdbController) startChecker() {
	// TODO 考虑在其他地方启动checker
	// 启动checker factory
	v1.CoreV1.CMDB().StartChecker()
	// 启动生产者
	go func() {
		timeout := 10 * time.Second
		// 从数据库中查询所有主机并进行检测
		for range time.Tick(timeout) {
			v1.CoreV1.CMDB().Host().StartHostCheck()
		}
	}()
}

func (c *cmdbController) initRoutes(ginEngine *gin.RouterGroup) {
	cmdbRoute := ginEngine.Group("/cmdb")
	{
		cmdbRoute.GET("/getHostGroupTree", c.GetHostGroupTree)
	}
	{
		cmdbRoute.GET("/pageHost", c.PageHost)
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
