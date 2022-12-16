package menu

import "github.com/gin-gonic/gin"

type menuController struct{}

func NewMenuRouter(ginEngine *gin.RouterGroup) {
	menu := menuController{}
	menu.initRoutes(ginEngine)
}

func (m *menuController) initRoutes(ginEngine *gin.RouterGroup) {
	menuRoute := ginEngine.Group("/menu")
	menu := &menuController{}
	{
		menuRoute.GET("/:authID/getMenuByAuthID", menu.GetMenusByAuthID)
		menuRoute.GET("/getBaseMenuTree", menu.GetBaseMenus)
		menuRoute.POST("/add_base_menu", menu.AddBaseMenu)
		menuRoute.POST("/add_menu_authority", menu.AddMenuAuthority)
	}
}
