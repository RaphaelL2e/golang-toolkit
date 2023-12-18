package redis

import (
	"github.com/go-redis/redis/v8"

	"github.com/RaphaelL2e/golang-toolkit/config"
	"github.com/RaphaelL2e/golang-toolkit/constant"
)

var RedisDilimiters = []string{"$", ":"}

type RedisClient struct {
	namespace string

	redis.UniversalClient
}

func NewRedisClient(config *config.RedisConfig) *RedisClient {
	if config.Type == "" {
		config.Type = constant.RedisStandaloneMode
	}

	if config.Type == constant.RedisClusterMode {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        config.ClusterAddress,
			Username:     config.Username,
			Password:     config.Password,
			DialTimeout:  constant.RedisDailTimeout,
			ReadTimeout:  constant.RedisReadTimeout,
			WriteTimeout: constant.RedisWriteTimeout,
		})
		return &RedisClient{
			UniversalClient: client,
			namespace:       config.Namespace,
		}
	} else {
		client := redis.NewClient(&redis.Options{
			Addr:         config.Address,
			Network:      config.Network,
			Username:     config.Username,
			Password:     config.Password,
			DB:           config.DB,
			DialTimeout:  constant.RedisDailTimeout,
			ReadTimeout:  constant.RedisReadTimeout,
			WriteTimeout: constant.RedisWriteTimeout,
		})
		return &RedisClient{
			UniversalClient: client,
			namespace:       config.Namespace,
		}
	}
}
