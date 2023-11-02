package config

import (
	"github.com/lizaiganshenmo/mixStew/config"
	"github.com/olivere/elastic/v7"
)

func GetMysqlDSN(srvName string) (string, error) {
	return config.GetMysqlDSN(ConfMap, srvName)
}

func GetEsClient(srvName string) (*elastic.Client, string, error) {
	return config.GetEsClient(ConfMap, srvName)
}
