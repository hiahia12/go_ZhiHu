package middleware

import (
	"github.com/gin-gonic/gin"
	"go_ZhiHu/app/global"
	"go_ZhiHu/utils/cookie"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		cookieConfig := global.Config.App.Cookie

		cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
			cookie.Option{
				Config: cookieConfig.Cookie,
				Ctx:    c,
			})
ok:=cookieWriter.
	}

}
