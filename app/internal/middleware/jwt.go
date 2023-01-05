package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"go_ZhiHu/app/global"
	"go_ZhiHu/app/internal/model/response"
	"go_ZhiHu/utils/cookie"
	myjwt "go_ZhiHu/utils/jwt"
	"net/http"
	"time"
)

var Id int64

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		cookieConfig := global.Config.App.Cookie

		cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
			cookie.Option{
				Config: cookieConfig.Cookie,
				Ctx:    c,
			})

		ok := cookieWriter.Get("x-token", &token)
		if token == "" || !ok {
			response.Fail(c, http.StatusUnauthorized, 1, "not logged in")
			c.Abort()
			return
		}

		jwtConfig := global.Config.Middleware.Jwt
		j := myjwt.NewJWT(&myjwt.Config{
			SecretKey: jwtConfig.SecretKey,
		})

		mc, err := j.ParseToken(token)
		fmt.Print(mc)
		if err != nil {
			response.Fail(c, http.StatusBadRequest, 1, err.Error())
			c.Abort()
			return
		}
		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime && mc.ExpiresAt.Unix()-time.Now().Unix() > 0 {
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Middleware.Jwt.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(mc)
			newClaims, _ := j.ParseToken(newToken)
			cookieWriter.Set("x-token", newToken)
			err = global.Rdb.Set(c,
				fmt.Sprintf("jwt；%d", newClaims.BaseClaims.Id),
				newToken,
				time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
			if err != nil {
				global.Logger.Error("set redis key failed.",
					zap.Error(err),
					zap.String("key", "jwt:[id]"), zap.Int64("id", newClaims.BaseClaims.Id),
				)
				response.InternalErr(c)
				c.Abort()
				return
			}
		}
		c.Set("id", mc.BaseClaims.Id)
		c.Next()
	}

}
