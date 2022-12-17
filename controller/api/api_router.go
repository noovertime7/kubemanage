package api

import "github.com/gin-gonic/gin"

type apiController struct{}

func NewApiRouter(ginGroup *gin.RouterGroup) {
	api := &apiController{}
	api.initRoutes(ginGroup)
}

func (a *apiController) initRoutes(ginGroup *gin.RouterGroup) {
	apiRoute := ginGroup.Group("/sysApi")
	apiRoute.GET("/getAPiList", a.GetApiList)
}
