package operation

import (
	"github.com/gin-gonic/gin"
)

type operationController struct{}

func NewOperationRouter(ginEngine *gin.RouterGroup) {
	opera := operationController{}
	opera.initRoutes(ginEngine)
}

func (o *operationController) initRoutes(ginEngine *gin.RouterGroup) {
	operaRoute := ginEngine.Group("/operation")
	opera := &operationController{}
	{
		operaRoute.GET("/get_operations", opera.GetOperationRecordList)
		operaRoute.DELETE("/:id/delete_operation", opera.DeleteOperationRecord)
		operaRoute.POST("/delete_operations", opera.DeleteOperationRecords)
	}
}
