package constants

const (
	// jwt
	SecretKey     = "secret key"
	IdentityKey   = "uid"
	TokenHeadName = "Token"

	// snowflake
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

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

	MixStewServiceName = "mixStewService"
)
