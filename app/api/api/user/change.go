package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/internal/model"
	service2 "go_ZhiHu/app/internal/service"
	"net/http"
)

type ChangApi struct {
}

var insChange = ChangApi{}

func (a *SignApi) ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if oldPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "oldPassword cannot be null",
			"ok":   false,
		})
		return
	}
	if newPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "newPassword cannot be null",
			"ok":   false,
		})
		return
	}

	userSubject := &model.UserSubject{
		Username: username,
		Password: service2.User().User().EncryptPassword(oldPassword),
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
	encryptedPassword := service2.User().User().EncryptPassword(newPassword)
	user := service2.User().User().GetUser(c, username)
	err1 := service2.User().User().ChangePassword(c, encryptedPassword, user.Id)
	if err != nil {
		fmt.Print(err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "change password successfully",
		"ok":   true,
	})
}

func (a *SignApi) ChangeUsername(c *gin.Context) {
	newUsername := c.PostForm("newUsername")
	oldUsername := c.PostForm("oldUsername")
	if newUsername == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "newUsername cannot be null",
			"ok":   false,
		})
		return
	}
	if oldUsername == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "oldUsername cannot be null",
			"ok":   false,
		})
		return
	}
	err := service2.User().User().CheckUserIsExist(c, newUsername)
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
	user := service2.User().User().GetUser(c, oldUsername)
	err1 := service2.User().User().ChangeUsername(c, newUsername, user.Id)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "change username successfully",
		"ok":   true,
	})
}
