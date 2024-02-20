package cache

import (
	config "github.com/lizaiganshenmo/mixStew/cmd/interaction/configs"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	initRedis("redis_mixStew")
}

func initRedis(srvName string) {
	var err error
	RedisClient, err = config.GetRedisClient(srvName)
	if err != nil {
		panic(err)
	}
}
