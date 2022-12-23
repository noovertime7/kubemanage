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
		userRoute.POST("/register", user.RegisterUser)
		userRoute.POST("/update", user.UpdateUser)
		userRoute.POST("/login", user.Login)
		userRoute.GET("/loginout", user.LoginOut)
		userRoute.GET("/getinfo", user.GetUserInfo)
		userRoute.POST("/:id/getPage", user.PageUsers)
		userRoute.POST("/:id/set_auth", user.SetUserAuthority)
		userRoute.DELETE("/:id/delete_user", user.DeleteUser)
		userRoute.POST("/delete_users", user.DeleteUsers)
		userRoute.POST("/:id/change_pwd", user.ChangePassword)
		userRoute.PUT("/:id/reset_pwd", user.ResetPassword)
		userRoute.PUT("/:id/:action/lockUser", user.LockUser)
	}
	{
		userRoute.GET("/deptTree", user.GetDepartmentTree)
		userRoute.POST("/getDeptByPage", user.GetDeptByPage)
		userRoute.GET("/:id/deptUsers", user.GetDepartmentUsers)
	}
}
