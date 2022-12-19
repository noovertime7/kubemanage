package user

import "github.com/gin-gonic/gin"

type userController struct{}

func NewUserRouter(ginEngine *gin.RouterGroup) {
	u := userController{}
	u.initRoutes(ginEngine)
}

func (u *userController) initRoutes(ginEngine *gin.RouterGroup) {
	userRoute := ginEngine.Group("/user")
	user := &userController{}
	{
		userRoute.POST("/login", user.Login)
		userRoute.GET("/loginout", user.LoginOut)
		userRoute.GET("/getinfo", user.GetUserInfo)
		userRoute.POST("/:id/getPage", user.PageUsers)
		userRoute.PUT("/:id/set_auth", user.SetUserAuthority)
		userRoute.DELETE("/:id/delete_user", user.DeleteUser)
		userRoute.POST("/:id/change_pwd", user.ChangePassword)
		userRoute.PUT("/:id/reset_pwd", user.ResetPassword)
	}
	{
		userRoute.GET("/deptTree", user.GetDepartmentTree)
		userRoute.GET("/:id/deptUsers", user.GetDepartmentUsers)
	}
}
