package app_grpc_server

import (
	"time"
)

type Config struct {
	port              string
	connectionTimeout time.Duration
	shutDownTimeout   time.Duration
}

func NewConfig(port string, connectionTimeout time.Duration, shutDownTimeout time.Duration) Config {
	return Config{
		port:              port,
		connectionTimeout: connectionTimeout,
		shutDownTimeout:   shutDownTimeout,
	}
}
