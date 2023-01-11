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

func (a *FavouriteApi) GetFavourites(c *gin.Context) {
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	userId, _ := strconv.ParseInt(userid, 10, 64)
	favourites, err := service2.User().User().GetFavourites(c, userId)
	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  favourites,
		"ok":   true,
	})
}
func (a *FavouriteApi) AddFavouriteAnswer(c *gin.Context) {
	userid := c.PostForm("userid")
	answerid := c.PostForm("answerid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	if answerid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "answerid cannot be nil",
			"ok":   false,
		})
	}
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	answerId, _ := strconv.ParseInt(answerid, 10, 64)
	favouriteanswer := &model.MyFavouriteAnswerSubject{
		Answerid:    answerId,
		Favouriteid: favouriteId,
		Userid:      userId,
	}

	err := service2.User().User().AddFavouriteAnswer(c, favouriteanswer)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add favourite answer successfully",
		"ok":   true,
	})
}

func (a *FavouriteApi) AddFavouriteQuestion(c *gin.Context) {
	userid := c.PostForm("userid")
	questionid := c.PostForm("questionid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
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
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	favouritequestion := &model.MyFavouriteQuestionSubject{
		Questionid:  questionId,
		Favouriteid: favouriteId,
		Userid:      userId,
	}

	err := service2.User().User().AddFavouriteQuestion(c, favouritequestion)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add favourite question successfully",
		"ok":   true,
	})
}

func (a *FavouriteApi) AddFavouriteArticle(c *gin.Context) {
	userid := c.PostForm("userid")
	articleid := c.PostForm("articleid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	if articleid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "questionid cannot be nil",
			"ok":   false,
		})
	}
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	articleId, _ := strconv.ParseInt(articleid, 10, 64)
	favouriteArticle := &model.MyFavouriteArticleSubject{
		Articleid:   articleId,
		Favouriteid: favouriteId,
		Userid:      userId,
	}

	err := service2.User().User().AddFavouriteArticle(c, favouriteArticle)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "add favourite article successfully",
		"ok":   true,
	})
}

func (a *FavouriteApi) CancelFavouriteQuestion(c *gin.Context) {
	userid := c.PostForm("userid")
	questionid := c.PostForm("questionid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
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
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	questionId, _ := strconv.ParseInt(questionid, 10, 64)

	favouritequestion, err1 := service2.User().User().GetFavouriteQuestion(c, userId, questionId, favouriteId)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	err := service2.User().User().CancelFavouriteQuestion(c, favouritequestion.Id)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel favourite question successfully",
		"ok":   true,
	})
}

func (a *FavouriteApi) CancelFavouriteArticle(c *gin.Context) {
	userid := c.PostForm("userid")
	articleid := c.PostForm("articleid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	if articleid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "articleid cannot be nil",
			"ok":   false,
		})
	}
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	articleId, _ := strconv.ParseInt(articleid, 10, 64)

	favouritearticle, err1 := service2.User().User().GetFavouriteArticle(c, userId, articleId, favouriteId)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	err := service2.User().User().CancelFavouriteArticle(c, favouritearticle.Id)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel favourite article successfully",
		"ok":   true,
	})
}

func (a *FavouriteApi) CancelFavouriteAnswer(c *gin.Context) {
	userid := c.PostForm("userid")
	answerid := c.PostForm("answerid")
	favouriteid := c.PostForm("favouriteid")
	if favouriteid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "favouriteid cannot be nil",
			"ok":   false,
		})
	}
	if answerid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "answerid cannot be nil",
			"ok":   false,
		})
	}
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}

	favouriteId, _ := strconv.ParseInt(favouriteid, 10, 64)
	userId, _ := strconv.ParseInt(userid, 10, 64)
	answerId, _ := strconv.ParseInt(answerid, 10, 64)

	favouriteanswer, err1 := service2.User().User().GetFavouriteAnswer(c, userId, answerId, favouriteId)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	err := service2.User().User().CancelFavouriteAnswer(c, favouriteanswer.Id)

	if err != nil {
		fmt.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel favourite answer successfully",
		"ok":   true,
	})
}
