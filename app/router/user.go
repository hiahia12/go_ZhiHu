package router

import (
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/api/api"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	userApi := api.User()
	{
		userRouter.POST("/register")
		userRouter.POST("/login")
	}
	return userRouter
}
func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/user")
	return userRouter
}