package main

import "go_ZhiHu/boot"

func main() {
	boot.LoggerSetup()
	boot.ViperSetup()
	boot.MysqlSetup()
	boot.RedisSetup()

}
