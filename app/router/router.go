package router

import (
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/global"
	"go_ZhiHu/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ZapLogger(global.Logger), middleware.ZapRecovery(global.Logger, true))

	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(publicGroup)
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{
		routerGroup.InitUserWriteRouter(privateGroup)
	}
	r.Run(":8080")
	global.Logger.Info("initialize routers successfully!")
	return r
}
