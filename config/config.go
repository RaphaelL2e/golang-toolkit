package config

import "time"

type Config struct {
	ServerConfig Server `mapstructure:"server"`
}

type Server struct {
	RunMode       string
	HttpUrl       string
	HttpPort      int
	HttpsOpen     bool
	CertFile      string
	KeyFile       string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	PromqlURL     string
	PromqlTimeout time.Duration
}

type RedisConfig struct {
	Network        string
	DB             int
	Address        string
	Username       string
	Password       string
	Type           string
	ClusterAddress []string
	Namespace      string
}
