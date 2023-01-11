package user

import (
	"github.com/gin-gonic/gin"
	service2 "go_ZhiHu/app/internal/service"
	"net/http"
	"strconv"
)

type LikeApi struct{}

var insLike = LikeApi{}

func (a *LikeApi) AddLikeAnswer(c *gin.Context) {
	answerid := c.PostForm("answerid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	answerId, _ := strconv.ParseInt(answerid, 10, 64)
	s := service2.User().User().AddLikeAnswer(c, answerId, userId) //写入数据库未完成
	if s == "like today" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "like today",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "like answer successfully",
		"ok":   true,
	})
}

func (a *LikeApi) AddLikeQuestion(c *gin.Context) {
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	s := service2.User().User().AddLikeQuestion(c, questionId, userId)
	if s == "like today" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "like today",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "like question successfully",
		"ok":   true,
	})
}

func (a *LikeApi) AddLikeArticle(c *gin.Context) {
	articleid := c.PostForm("articleid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	articleId, _ := strconv.ParseInt(articleid, 10, 64)
	s := service2.User().User().AddLikeArticle(c, articleId, userId)
	if s == "like today" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "like today",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "like article successfully",
		"ok":   true,
	})
}

func (a *LikeApi) AddLikeComment(c *gin.Context) {
	commentid := c.PostForm("commentid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if commentid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "commentid cannot be nil",
			"ok":   false,
		})
	}

	userId, _ := strconv.ParseInt(userid, 10, 64)
	commentId, _ := strconv.ParseInt(commentid, 10, 64)
	s := service2.User().User().AddLikeComment(c, commentId, userId)
	if s == "like today" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "like today",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "like comment successfully",
		"ok":   true,
	})
}

func (a *LikeApi) CancelLikeAnswer(c *gin.Context) {
	answerid := c.PostForm("answerid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	answerId, _ := strconv.ParseInt(answerid, 10, 64)
	s := service2.User().User().CancelLikeAnswer(c, answerId, userId)
	if s == "answer didn't be liked" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "answer didn't be liked",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "like answer successfully",
		"ok":   true,
	})
}

func (a *LikeApi) CancelLikeQuestion(c *gin.Context) {
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	s := service2.User().User().CancelLikeQuestion(c, questionId, userId)
	if s == "question didn't be liked " {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "question didn't be liked ",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel like question successfully",
		"ok":   true,
	})
}

func (a *LikeApi) CancelLikeArticle(c *gin.Context) {
	articleid := c.PostForm("articleid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
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

	userId, _ := strconv.ParseInt(userid, 10, 64)
	articleId, _ := strconv.ParseInt(articleid, 10, 64)
	s := service2.User().User().CancelLikeArticle(c, articleId, userId)
	if s == "article didn't be liked" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "article didn't be liked",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  " cancel like article successfully",
		"ok":   true,
	})
}

func (a *LikeApi) CancelLikeComment(c *gin.Context) {
	commentid := c.PostForm("commentid")
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "userid cannot be nil",
			"ok":   false,
		})
	}
	if commentid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "commentid cannot be nil",
			"ok":   false,
		})
	}

	userId, _ := strconv.ParseInt(userid, 10, 64)
	commentId, _ := strconv.ParseInt(commentid, 10, 64)
	s := service2.User().User().CancelLikeComment(c, commentId, userId)
	if s == "like today" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "comment didn't be liked ",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "cancel like comment successfully",
		"ok":   true,
	})
}
