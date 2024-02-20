package constants

const (
	// jwt
	SecretKey     = "secret key"
	IdentityKey   = "uid"
	TokenHeadName = "Token"

	// request id key
	RequestIdKey = "log_id"

	// jaeger collector add  -> accept OpenTelemetry Protocol (OTLP) over gRPC
	JaegerColAdd = "123.56.67.72:4317"
	JaegerAdd    = "123.56.67.72:6831"

	// snowflake
	SnowflakeWorkerID                = 0
	SnowflakeDatacenterID            = 0
	SnowflakeArticleWorkerID         = 1
	SnowflakeArticleDatacenterID     = 1
	SnowflakeInteractionWorkerID     = 2
	SnowflakeInteractionDatacenterID = 2

	// etcd add
	EtcdAdd = "123.56.67.72:2379"

	// Service
	UserServiceName = "userService"
	// listenADD
	UserServiceListenADD = "127.0.0.1:8001"
	// service limit
	MaxConnections = 1000
	MaxQPS         = 100

	FollowServiceName      = "followService"
	FollowServiceListenADD = "127.0.0.1:8002"

	MixStewServiceName      = "mixStewService"
	ArticleServiceName      = "articleService"
	ArticleServiceListenADD = "127.0.0.1:8003"

	InteractionServiceName      = "interactionService"
	InteractionServiceListenADD = "127.0.0.1:8004"
)
