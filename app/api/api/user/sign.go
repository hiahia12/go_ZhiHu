package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/global"
	"go_ZhiHu/app/internal/model"
	service2 "go_ZhiHu/app/internal/service"
	"go_ZhiHu/utils/cookie"
	"net/http"
)

type SignApi struct {
}

var insSign = SignApi{}

func (a *SignApi) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	err := service2.User().User().CheckUserIsExist(c, username)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		} else if err.Error() == "username already exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	userSubject := &model.UserSubject{}

	encryptedPassword := service2.User().User().EncryptPassword(password)

	userSubject.Username = username
	userSubject.Password = encryptedPassword
	service2.User().User().CreateUser(c, userSubject)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "register successfully",
		"ok":   true,
	})
}

func (a *SignApi) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	userSubject := &model.UserSubject{
		Username: username,
		Password: service2.User().User().EncryptPassword(password),
	}

	err := service2.User().User().CheckPassword(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "invalid username or password":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

		return
	}

	tokenString, err := service2.User().User().GenerateToken(c, userSubject)
	fmt.Print(tokenString)
	if err != nil {

		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

	}

	cookieConfig := global.Config.App.Cookie

	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
		cookie.Option{
			Config: cookieConfig.Cookie,
			Ctx:    c,
		})

	cookieWriter.Set("x-token", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "login successfully",
		"ok":   true,
	})
}
