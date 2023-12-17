package router

import (
	"account-server/global"
	"account-server/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IFnRegistRoute = func(rgPubulic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouter []IFnRegistRoute
)

// RegistRoute 注册路由回电
func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRouter = append(gfnRouter, fn)
}

func InitBasePlatFormRoutes() {
	InitUserRoutes()
	InitLoginRoutes()
}

// InitRouter 初始化路由
func InitRouter() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()
	r.Use(middleware.Cors())
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	rgPubulic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1/")
	rgAuth.Use(middleware.JWTAuthMiddleware())

	InitBasePlatFormRoutes()
	for _, fnRegistRoute := range gfnRouter {
		fnRegistRoute(rgPubulic, rgAuth)
	}
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8088"
	}

	server := &http.Server{Addr: fmt.Sprintf(":%s", stPort), Handler: r}

	go func() {
		global.Logger.Info(fmt.Sprintf("Star listen:%s", stPort))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Star Server Error:%s", err.Error()))
		}
	}()
	<-ctx.Done()

	ctx, cancelShutdowm := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdowm()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error:%s", err.Error()))
	}
	global.Logger.Info("Stop Server Success")
}
