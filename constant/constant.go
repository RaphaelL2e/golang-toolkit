package constant

import "time"

const (
	HttpTimeout = 1 * time.Second
	GrpcTimeout = 1 * time.Second

	GrpcNetwork = "tcp"

	RedisStandaloneMode = "standalone"
	RedisClusterMode    = "cluster"
	RedisDailTimeout    = 1 * time.Second
	RedisReadTimeout    = 1 * time.Second
	RedisWriteTimeout   = 1 * time.Second

	RedsyncTries  = 1
	RedsyncExpiry = 30 * time.Second
)
