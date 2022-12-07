package main

import "go_ZhiHu/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlSetup()
	boot.RedisSetup()
	boot.ServerSetup()
}
