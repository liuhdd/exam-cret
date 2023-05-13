package config

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	conf := GetConfig()
	once.Do(func() {
		conf := conf.Sub("redis")
		redisClient = redis.NewClient(&redis.Options{
			Addr:     conf.GetString("addr"),
			Password: conf.GetString("password"),
			DB:       conf.GetInt("db"),
		})
		ctx := context.Background()
		redisClient.XGroupCreate(ctx, "action", "g1", "$").Result()
	})

	return redisClient
}
