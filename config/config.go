package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

func InitConfig() *Config {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.ReadInConfig()
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		zap.S().Fatalln("failed to unmarshal config", err)
		return nil
	}
	zap.S().Infof("config: %+v", cfg)
	return &cfg
}
