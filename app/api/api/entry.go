package api

import "go_ZhiHu/app/api/api/user"

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
