package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/olivere/elastic/v7"
)

// get mysqlDSN by map conf
func GetMysqlDSN(confMap map[string]interface{}, srvName string) (string, error) {
	val, ok := confMap[srvName]
	if !ok {
		return "", errors.New(fmt.Sprintf("unkonwn srvname: %s", srvName))
	}

	var mc MySQLConf
	err := mapstructure.Decode(val, &mc)
	if err != nil {
		return "", err
	}

	dsn := strings.Join([]string{mc.MySQL.Username, ":",
		mc.MySQL.Password, "@tcp(", mc.MySQL.Addr, ")/", mc.MySQL.Database, "?charset=" + mc.MySQL.Charset + "&parseTime=true"}, "")
	return dsn, nil
}

// get esClient by map conf
func GetEsClient(confMap map[string]interface{}, srvName string) (*elastic.Client, string, error) {
	val, ok := confMap[srvName]
	if !ok {
		return nil, "", errors.New(fmt.Sprintf("unkonwn srvname: %s", srvName))
	}

	var ec EsConf
	err := mapstructure.Decode(val, &ec)
	if err != nil {
		return nil, "", err
	}

	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", ec.ES.Host, ec.ES.Port)),
		elastic.SetBasicAuth(ec.ES.UserName, ec.ES.Password),
		elastic.SetSniff(false),
	)

	return client, ec.ES.Host, err

}
