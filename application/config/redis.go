package config

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once

var cilent *redis.Client

func GetRedisCilent() *redis.Client{

	once.Do(func() {
		conf := v.Sub("redis")
		cilent = redis.NewClient(&redis.Options{
			Addr:     conf.GetString("addr"),
			Password: conf.GetString("password"), 
			DB:       conf.GetInt("db"),
		})
		
	})

	return cilent
}

