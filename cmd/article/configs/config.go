package config

import (
	"github.com/lizaiganshenmo/mixStew/config"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
)

func GetMysqlDSN(srvName string) (string, error) {
	return config.GetMysqlDSN(confMap, srvName)
}

func GetEsClient(srvName string) (*elastic.Client, string, error) {
	return config.GetEsClient(confMap, srvName)
}

func GetRedisClient(srvName string) (*redis.Client, error) {
	return config.GetRedisClient(confMap, srvName)
}
