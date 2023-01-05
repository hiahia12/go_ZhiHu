package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go_ZhiHu/utils/cookie"
	"time"
)

type CustomClaims struct {
	BufferTime int64
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	Id         int64
	Username   string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetClaims(secret string, cookie *cookie.Cookie) (*CustomClaims, error) {
	var token string
	ok := cookie.Get("x-token", &token)

	//token, err := c.Cookie("x-token")
	if !ok {
		err := errors.New("get token by cookie failed")
		return nil, err
	}
	j := NewJWT(&Config{SecretKey: secret})
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Print(err)
		err := errors.New("parse token failed")
		return nil, err
	}
	return claims, nil
}

func GetUserInfo(secret string, cookie *cookie.Cookie) (*BaseClaims, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return nil, err
	} else {
		return &cl.BaseClaims, nil
	}
}

func GetUserID(secret string, cookie *cookie.Cookie) (int64, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return -1, err
	} else {
		return cl.BaseClaims.Id, nil
	}
}

//func GetUserID(c *gin.Context) (int64, error) {
//	var cookie1 = &cookie.Cookie{}
//	cookie := global.Config.App.Cookie
//	cookies, err := c.Cookie("x-token")
//	val := cookies
//	if err != nil {
//		return -1, err
//	}
//	parts := strings.SplitN(val, "|", 3)
//	if len(parts) != 3 {
//		return -1, err
//	}
//	vs := parts[0]
//	timestamp := parts[1]
//	sig := parts[2]
//	h := hmac.New(sha256.New, []byte(cookie.Secret))
//	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)
//
//	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
//		return -1, err
//	}
//	res, _ := base64.URLEncoding.DecodeString(vs)
//	tempDate := res
//	_ = json.Unmarshal([]byte(tempDate), cookie1)
//	jwtConfig := global.Config.Middleware.Jwt
//	j := NewJWT(&Config{
//		SecretKey: jwtConfig.SecretKey,
//	})
//	claims, err := j.ParseToken(cookies)
//	if err != nil {
//		fmt.Print(err)
//		return -1, err
//	}
//	fmt.Print(cookies)
//	fmt.Print(claims)
//	return claims.Id, err
//}
