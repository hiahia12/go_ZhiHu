package service

import "go_ZhiHu/app/internal/service/user"

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
