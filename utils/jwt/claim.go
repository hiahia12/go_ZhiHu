package jwt

import (
	"github.com/golang-jwt/jwt/v4"
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
