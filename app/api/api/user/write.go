package user

import (
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/internal/model"
	service2 "go_ZhiHu/app/internal/service"
	"net/http"
	"strconv"
)

type WriteApi struct {
}

var insWrite = WriteApi{}

func (a *WriteApi) GetQuestions(c *gin.Context) {
	questions := service2.User().User().GetQuestions(c)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  questions,
		"ok":   true,
	})
}

func (a *WriteApi) GetAnswer(c *gin.Context) {
	questionid := c.PostForm("questionid")
	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	answers := service2.User().User().GetAnswer(c, questionId)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  answers,
		"ok":   true,
	})
}

func (a *WriteApi) GetComment(c *gin.Context) {
	answerId := c.PostForm("answerid")
	answerid, _ := strconv.ParseInt(answerId, 10, 64)
	comments := service2.User().User().GetComment(c, answerid)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  comments,
		"ok":   true,
	})
}

func (a *WriteApi) GetArticle(c *gin.Context) {
	articles := service2.User().User().GetArticles(c)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  articles,
		"ok":   true,
	})
}

func (a *WriteApi) WriteQuestion(c *gin.Context) {
	question := c.PostForm("question")
	//userid, err := jwt.GetUserID(c)
	//fmt.Print(userid)
	//if err != nil {
	//	fmt.Print(err)
	//	return
	//}
	////cookieConfig := global.Config.App.Cookie
	//cookieWriter := cookie2.NewCookieWriter(cookieConfig.Secret,
	//	cookie2.Option{
	//		Config: cookieConfig.Cookie,
	//		Ctx:    c,
	//	})
	//
	//cookie.Opt.Config.Value = cookies
	//cookie.Opt.Config.Name = "x-token"
	//cookie.Opt.Config.Path = "/"
	//cookie.Opt.Config.HttpOnly = cookieWriter.Opt.Config.HttpOnly
	//cookie.Opt.Config.Domain = cookieWriter.Opt.Config.Domain
	//cookie.Opt.Config.MaxAge = cookieWriter.Opt.Config.MaxAge
	//cookie.Opt.Config.Secure = cookieWriter.Opt.Config.Secure
	////cookie.Opt.Config.SameSite = http.SameSite(1)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": http.StatusBadRequest,
	//		"msg":  err,
	//		"ok":   false,
	//	})
	//	return
	//}
	if question == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "question cannot be nil",
			"ok":   false,
		})
		return
	}

	questionSubject := &model.Question{
		Question: question,
		Askerid:  1,
	}
	service2.User().User().WriteQuestion(c, questionSubject)

	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusOK,
		"msg":  "write question successfully",
		"ok":   true,
	})
}

func (a *WriteApi) WriteAnswer(c *gin.Context) {
	answer := c.PostForm("answer")
	questionid := c.PostForm("questionid")
	username := c.PostForm("username")
	if answer == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "answer cannot be nil",
			"ok":   false,
		})
		return
	}
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be nil",
			"ok":   false,
		})
		return
	}
	if questionid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "questionid cannot be nil",
			"ok":   false,
		})
		return
	}

	questionId, _ := strconv.ParseInt(questionid, 10, 64)
	err := service2.User().User().CheckQuestionIsExist(c, questionId)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		} else if err.Error() == "answer not found" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	user := service2.User().User().GetUser(c, username)
	answersubject := &model.AnswerSubject{
		Answer:     answer,
		Writerid:   user.Id,
		Questionid: questionId,
	}
	service2.User().User().WriteAnswer(c, answersubject)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "write answer successfully",
		"ok":   true,
	})
}

func (a *WriteApi) WriteComment(c *gin.Context) {
	username := c.PostForm("username")
	comment := c.PostForm("comment")
	answerid := c.PostForm("answerid")
	if comment == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "comment cannot be nil",
			"ok":   false,
		})
		return
	}
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be nil",
			"ok":   false,
		})
		return
	}
	if answerid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "answerid cannot be nil",
			"ok":   false,
		})
		return
	}
	answerId, _ := strconv.ParseInt(answerid, 10, 64)
	err := service2.User().User().CheckAnswerIsExist(c, answerId)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		} else if err.Error() == "answer not found" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	user := service2.User().User().GetUser(c, username)

	commentsubject := &model.Comment{
		Comment:  comment,
		Answerid: answerId,
		Writerid: user.Id,
	}
	service2.User().User().WriteComment(c, commentsubject)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "write comment successfully",
		"ok":   true,
	})
}

func (a *WriteApi) WriteArticle(c *gin.Context) {
	article := c.PostForm("article")
	username := c.PostForm("username")
	if article == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "article cannot be nil",
			"ok":   false,
		})
		return
	}
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be nil",
			"ok":   false,
		})
		return
	}
	user := service2.User().User().GetUser(c, username)

	articleSubject := &model.ArticleSubject{
		Article:  article,
		Writerid: user.Id,
	}

	service2.User().User().WriteArticle(c, articleSubject)

	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusOK,
		"msg":  "write article successfully",
		"ok":   true,
	})
}
