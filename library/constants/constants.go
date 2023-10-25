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
	EtcdAdd = "47.93.215.110:2379"

	// Service
	UserServiceName = "userService"
	// listenADD
	ListenADD = "127.0.0.1:8001"
	// service limit
	MaxConnections = 1000
	MaxQPS         = 100
)
