package app_http_server

import "time"

type ServerConfig struct {
	addr            string
	timeOut         time.Duration
	shutdownTimeout time.Duration
}

func NewHttpServerConfig() ServerConfig {
	return ServerConfig{
		addr:            ":8081",
		timeOut:         10 * time.Second,
		shutdownTimeout: 15 * time.Second,
	}
}
