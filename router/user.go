package router

import (
	"account-server/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPubulic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPubulic.POST("login", userApi.Login)
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("/info/:id", func(c *gin.Context) {
				api.ResponseSuccess(c, 1)
			})
			rgAuthUser.POST("/addUser", userApi.AddUser)
			rgAuthUser.GET("/getUser/:id", userApi.GetUserById)
			rgAuthUser.POST("/list", userApi.GetUserList)
			rgAuthUser.POST("/edit", userApi.EditUser)
			rgAuthUser.GET("/del/:id", userApi.DelUser)
		}
	})
}
