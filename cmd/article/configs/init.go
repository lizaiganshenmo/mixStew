package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

var (
	confMap = map[string]interface{}{} // 存储所有配置文件信息
)

// 本地静态配置文件加载
func staicConfInit(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var staticViper *viper.Viper
		if !info.IsDir() && filepath.Ext(path) == ".toml" {
			staticViper = viper.New()
			staticViper.SetConfigFile(path)

			err := staticViper.ReadInConfig()
			if err != nil {
				return err
			}

			conf := staticViper.AllSettings()
			val, ok := conf["service_name"]
			if !ok {
				return errors.New(fmt.Sprintf("wrong file conf format: %s\n", path))
			}

			confMap[val.(string)] = conf

		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}

func dynamicConfInit(path string, srvName string) {

}

func Init() {
	staicConfInit("./configs")
}
