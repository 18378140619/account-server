package router

import (
	"account-server/api"
	"github.com/gin-gonic/gin"
)

func InitLoginRoutes() {
	RegistRoute(func(rgPubulic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		loginApi := api.NewLoginApi()
		rgPubulic.POST("backOpenid", loginApi.WxCodeGetOpenID)
		rgAuthUser := rgAuth.Group("wx")
		{
			rgAuthUser.GET("/del/:id", loginApi.WxCodeGetOpenID)
		}
	})
}
