package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/internal/model"
	service2 "go_ZhiHu/app/internal/service"
	"net/http"
	"strconv"
)

type FavouriteApi struct {
}

var insFavourite = FavouriteApi{}

func (a *FavouriteApi) CreatFavourites(c *gin.Context) {
	userid := c.PostForm("userid")
	public := c.PostForm("public")
	name := c.PostForm("name")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if public == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "public cannot be nil",
			"ok":   false,
		})
	}
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "name cannot be nil",
			"ok":   false,
		})
	}

	publics, _ := strconv.ParseInt(public, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)

	favourites := &model.FavouriteSubject{
		Userid:          userId,
		FavouriteNumber: 0,
		Public:          int(publics),
		Name:            name,
	}

	err := service2.User().User().CreatFavourites(c, favourites)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "creat favourites successful",
		"ok":   true,
	})

}

func (a *FavouriteApi) DeleteFavourites(c *gin.Context) {
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	err := service2.User().User().DeleteFavourites(c, favouriteId)
	if err != nil {
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "delete favourite successful",
		"ok":   true,
	})
}

func (a *FollowApi) AddFavouriteAnswer(c *gin.Context) {

}

func (a *FollowApi) AddFavouriteQuestion(c *gin.Context) {

}

func (a *FollowApi) AddFavouriteArticle(c *gin.Context) {

}

func (a *FollowApi) CancelFavouriteArticle(c *gin.Context) {

}

func (a *FollowApi) CancelFavouriteQuestion(c *gin.Context) {

}

func (a *FollowApi) CancelFavouriteAnswer(c *gin.Context) {

}
