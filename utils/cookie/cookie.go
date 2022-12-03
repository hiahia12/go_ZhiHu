package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	Cookie struct {
		Secret string
		Opt    Option
	}
	Option struct {
		Config http.Cookie
		Ctx    *gin.Context
	}

	Config struct {
		Secret string
		http.Cookie
	}
)

func NewCookieWriter(secret string, opt ...Option) *Cookie {
	if len(opt) == 0 {
		return &Cookie{Secret: secret}
	} else {
		return &Cookie{
			Secret: secret,
			Opt:    opt[0],
		}
	}

}
func setSecureCookie(c *Cookie, name, value string) {
	vs := base64.URLEncoding.EncodeToString([]byte(value))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha256.New, []byte(c.Secret))
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)

	sig := fmt.Sprintf("%02x", h.Sum(nil))
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")

	http.SetCookie(c.Opt.Ctx.Writer, &http.Cookie{
		Name:     name,
		Value:    cookie,
		Path:     "/",
		Domain:   c.Opt.Config.Domain,
		MaxAge:   c.Opt.Config.MaxAge,
		Secure:   c.Opt.Config.Secure,
		HttpOnly: c.Opt.Config.HttpOnly,
		SameSite: http.SameSite(1),
	})
}
func (c *Cookie) Set(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}
