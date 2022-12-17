package system

import (
	"github.com/gin-gonic/gin"
)

type systemController struct{}

func NewSystemController(ginEngine *gin.RouterGroup) {
	sys := systemController{}
	sys.initRoutes(ginEngine)
}

func (s *systemController) initRoutes(ginEngine *gin.RouterGroup) {
	sysRoute := ginEngine.Group("/system", s.GetSystemState)
	sysRoute.GET("state")
}
