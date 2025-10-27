package global

import "bluebell/config"

var Config *config.Config

func InitConfig() {
	Config = config.InitConfig()
}
