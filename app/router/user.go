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
		userRouter.GET("/getAnswer", userApi.Write().GetAnswer)
		userRouter.GET("/getComment", userApi.Write().GetComment)

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
		userRouter.POST("/writeArticle", userApi.Write().WriteArticle)
		userRouter.POST("/writeQuestion", userApi.Write().WriteQuestion)
		userRouter.POST("/writeAnswer", userApi.Write().WriteAnswer)
		userRouter.POST("/writeComment", userApi.Write().WriteComment)

		userRouter.POST("/AddLikeArticle", userApi.Like().AddLikeArticle)
		userRouter.POST("/AddLikeQuestion", userApi.Like().AddLikeQuestion)
		userRouter.POST("/AddLikeAnswer", userApi.Like().AddLikeAnswer)
		userRouter.POST("/AddLikeComment", userApi.Like().AddLikeComment)
		userRouter.POST("/CancelLikeArticle", userApi.Like().CancelLikeArticle)
		userRouter.POST("/CancelLikeQuestion", userApi.Like().CancelLikeQuestion)
		userRouter.POST("/CancelLikeAnswer", userApi.Like().CancelLikeAnswer)
		userRouter.POST("/CancelLikeComment", userApi.Like().CancelLikeComment)

		userRouter.POST("/AddFollowUser", userApi.Follow().AddFollowUser)
		userRouter.POST("/AddFollowQuestion", userApi.Follow().AddFollowQuestion)
		userRouter.POST("/AddFollowFavourite", userApi.Follow().AddFollowFavourite)
		userRouter.POST("/CancelFollowUser", userApi.Follow().CancelFollowUser)
		userRouter.POST("/CancelFollowQuestion", userApi.Follow().CancelFollowQuestion)
		userRouter.POST("/CancelFollowFavourite", userApi.Follow().CancelFollowFavourite)

		userRouter.POST("/GetFavourites", userApi.Favourite().GetFavourites)
		userRouter.POST("/CreatFavourites", userApi.Favourite().CreatFavourites)
		userRouter.POST("/DeleteFavourites", userApi.Favourite().DeleteFavourites)
		userRouter.POST("/AddFavouriteAnswer", userApi.Favourite().AddFavouriteAnswer)
		userRouter.POST("/AddFavouriteQuestion", userApi.Favourite().AddFavouriteQuestion)
		userRouter.POST("/AddFavouriteArticle", userApi.Favourite().AddFavouriteArticle)
		userRouter.POST("/CancelFavouriteAnswer", userApi.Favourite().CancelFavouriteAnswer)
		userRouter.POST("/CancelFavouriteQuestion", userApi.Favourite().CancelFavouriteQuestion)
		userRouter.POST("/CancelFavouriteArticle", userApi.Favourite().CancelFavouriteArticle)

		userRouter.POST("/ChangeUsername", userApi.Change().ChangeUsername)
		userRouter.POST("/ChangePassword", userApi.Change().ChangePassword)

	}
	return userRouter
}
