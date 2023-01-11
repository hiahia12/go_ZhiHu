package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/internal/model"
	service2 "go_ZhiHu/app/internal/service"
	"net/http"
	"strconv"
)

type FollowApi struct {
}

var insFollow = FollowApi{}

func (a *FollowApi) GetFollowQuestions(c *gin.Context) {

}

func (a *FollowApi) GetFollowFavourites(c *gin.Context) {

}

func (a *FollowApi) AddFollowQuestion(c *gin.Context) {
	questionid := c.PostForm("questionid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if questionid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "questionid cannot be nil",
			"ok":   false,
		})
	}
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followquestion := &model.FollowQuestionSubject{
		Userid:     userId,
		Questionid: questionId,
	}
	err := service2.User().User().AddFollowQuestion(c, followquestion)
	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add follow question successful",
		"ok":   true,
	})
}

func (a *FollowApi) AddFollowFavourite(c *gin.Context) {
	userid := c.PostForm("userid")
	favouriteid := c.PostForm("favouriteid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followfavourite := &model.FollowFavouriteSubject{
		Userid:      userId,
		Favouriteid: favouriteId,
	}
	err := service2.User().User().AddFollowFavourite(c, followfavourite)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add follow favourite successful",
		"ok":   true,
	})
}

func (a *FollowApi) AddFollowUser(c *gin.Context) {
	userid := c.PostForm("userid")
	followuserid := c.PostForm("followuserid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if followuserid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "followuserid cannot be nil",
			"ok":   false,
		})
	}
	followuserId, _ := strconv.ParseInt(followuserid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followuser := &model.FollowUserSubject{
		Userid:       userId,
		FollowUserid: followuserId,
	}
	err := service2.User().User().AddFollowUser(c, followuser)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add follow user successful",
		"ok":   true,
	})
}

func (a *FollowApi) CancelFollowQuestion(c *gin.Context) {
	questionid := c.PostForm("questionid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if questionid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "questionid cannot be nil",
			"ok":   false,
		})
	}
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followquestion := &model.FollowQuestionSubject{
		Userid:     userId,
		Questionid: questionId,
	}
	err := service2.User().User().CancelFollowQuestion(c, followquestion)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel follow question successful",
		"ok":   true,
	})
}

func (a *FollowApi) CancelFollowFavourite(c *gin.Context) {
	userid := c.PostForm("userid")
	favouriteid := c.PostForm("favouriteid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followfavourite := &model.FollowFavouriteSubject{
		Userid:      userId,
		Favouriteid: favouriteId,
	}
	err := service2.User().User().CancelFollowFavourite(c, followfavourite)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel follow favourite successful",
		"ok":   true,
	})
}

func (a *FollowApi) CancelFollowUser(c *gin.Context) {
	userid := c.PostForm("userid")
	followuserid := c.PostForm("followuserid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if followuserid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "followuserid cannot be nil",
			"ok":   false,
		})
	}
	followuserId, _ := strconv.ParseInt(followuserid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	followuser := &model.FollowUserSubject{
		Userid:       userId,
		FollowUserid: followuserId,
	}
	err := service2.User().User().CancelFollowUser(c, followuser)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel follow user successful",
		"ok":   true,
	})
}
