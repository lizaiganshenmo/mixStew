package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func GetMysqlDSN(srvName string) (string, error) {
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
