package main

import (
	"bluebell/config"
	"bluebell/global"
	"bluebell/pkg/snowflake"
	"bluebell/web/router"
)

func init() {
	global.InitLogger()
	global.Config = config.InitConfig()
	global.InitMysql()
	global.InitRedis()
	global.InitProducer()
	global.InitTranslation()
	snowflake.Init("2025-01-01", 1)
}

func main() {

	r := router.InitRouter()
	r.Run(":9090")

	defer global.Close()
}
