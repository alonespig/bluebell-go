package global

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Redis *redis.Client

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
