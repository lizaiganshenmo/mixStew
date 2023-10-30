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
	// runtime_viper.SetConfigType("yaml")
	// etcdAddr := os.Getenv("ETCD_ADDR")

	// if etcdAddr == "" {
	// 	panic(errors.New("not found etcd addr in env"))
	// }

	// Etcd = &etcd{Addr: etcdAddr}

	// // use etcd for config save
	// err := runtime_viper.AddRemoteProvider("etcd3", Etcd.Addr, "/config/config.yaml")

	// if err != nil {
	// 	panic(err)
	// }

	// klog.Infof("config path: %v\n", path)

	// if err := runtime_viper.ReadRemoteConfig(); err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		klog.Fatal("could not find config files")
	// 	} else {
	// 		klog.Fatal("read config error: %v", err)
	// 	}
	// 	klog.Fatal(err)
	// }

	// configMapping(service)

	// klog.Infof("all keys: %v\n", runtime_viper.AllKeys())

	// // 持续监听配置
	// runtime_viper.OnConfigChange(func(e fsnotify.Event) {
	// 	klog.Infof("config file changed: %v\n", e.String())
	// })
	// runtime_viper.WatchConfig()
}
