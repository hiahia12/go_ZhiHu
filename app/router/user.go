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
		userRouter.POST("/register", userApi.Sign().Register)
		userRouter.POST("/login", userApi.Sign().Login)
		userRouter.GET("/getQuestion", userApi.Write().GetQuestions)
	}
	return userRouter
}
func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/user")
	return userRouter
}

func (r *UserRouter) InitUserWriteRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	userApi := api.User()
	{
		userRouter.POST("/writeQuestion", userApi.Write().WriteQuestion)

	}
	return userRouter
}
