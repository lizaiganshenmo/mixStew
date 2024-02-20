package resource

import (
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	config "github.com/lizaiganshenmo/mixStew/cmd/interaction/configs"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	eslogrus "github.com/lizaiganshenmo/mixStew/library/logrus"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// globbar var
var (
	ESClient *elastic.Client
	ESHost   string
)

func Init() {
	ESInit("es_mixStew")

	logInit()
}

func ESInit(srvName string) {
	var err error
	ESClient, ESHost, err = config.GetEsClient(srvName)
	if err != nil {
		panic(err)
	}
}

func logInit() {
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(esHookLog())))
}

func esHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(ESClient, ESHost, logrus.DebugLevel, strings.ToLower(constants.InteractionServiceName))
	if err != nil {
		panic(err)
	}

	return hook
}
