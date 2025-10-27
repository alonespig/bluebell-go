package main

import (
	"bluebell/config"
	"bluebell/global"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
)

func init() {
	logger.InitLogger()
	global.Config = config.InitConfig()
	global.InitMysql()
	global.InitRedis()
	global.InitProducer()
	snowflake.Init("2025-01-01", 1)

}

func main() {

	r := router.InitRouter()
	r.Run(":9090")

	defer global.Close()
}
