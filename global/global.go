package global

import (
	"bluebell/config"
	"bluebell/model"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client
var Config *config.Config

func InitMysql() {
	cfg := Config.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalln("failed to connect database", err)
		return
	}
	DB = db
	DB.AutoMigrate(&model.User{},
		&model.Community{},
		&model.Post{})
}

func InitRedis() {
	cfg := Config.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		zap.S().Fatalln("failed to connect redis", err)
		return
	}
}
