package config

type ConfMeta struct {
	ServiceName string `mapstructure:"service-name"`
}

type Server struct {
	ConfMeta
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID      int64 `mapstructure:"worker-id"`
	DatancenterID int64 `mapstructure:"datancenter-id"`
}

type service struct {
	ConfMeta
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}

type MySQLConf struct {
	ConfMeta
	MySQL struct {
		Addr     string
		Database string
		Username string
		Password string
		Charset  string
	}
}

type jaeger struct {
	Addr string
}

type etcd struct {
	Addr string
}

type rabbitMQ struct {
	Addr     string
	Username string
	Password string
}

type redis struct {
	Addr     string
	Password string
}

type oss struct {
	Endpoint        string
	AccessKeyID     string `mapstructure:"accessKey-id"`
	AccessKeySecret string `mapstructure:"accessKey-secret"`
	BucketName      string
	MainDirectory   string `mapstructure:"main-directory"`
}

type elasticsearch struct {
	Addr string
	Host string
}

// type config struct {
// 	Server        Server
// 	Snowflake     snowflake
// 	MySQL         MySQLConf
// 	Jaeger        jaeger
// 	Etcd          etcd
// 	RabbitMQ      rabbitMQ
// 	Redis         redis
// 	OSS           oss
// 	Elasticsearch elasticsearch
// }
