package app_repository

import (
	"time"
)

type Config struct {
	User            string        `env: User`
	Password        string        `env: Password`
	Host            string        `env: Host`
	Port            string        `env: Port`
	DataBase        string        `env: DataBase`
	MaxConns        int           `env: MaxConns`
	MinConns        int           `env: MinConns`
	MaxConnLifeTime time.Duration `env: MaxConnLifeTime`
	MaxConnIdleTime time.Duration `env: MaxConnIdleTime`
}

var kindaEnvFileData = Config{
	User:            "Dan",
	Password:        "DanDan",
	Host:            "localhost",
	Port:            "8080",
	DataBase:        "articlesAuth",
	MaxConns:        25,
	MinConns:        5,
	MaxConnLifeTime: 5 * time.Minute,
	MaxConnIdleTime: 2 * time.Minute,
}

func NewConfig() Config {
	return Config{
		User:            kindaEnvFileData.User,
		Password:        kindaEnvFileData.Password,
		Host:            kindaEnvFileData.Host,
		Port:            kindaEnvFileData.Port,
		DataBase:        kindaEnvFileData.DataBase,
		MaxConns:        kindaEnvFileData.MaxConns,
		MinConns:        kindaEnvFileData.MinConns,
		MaxConnLifeTime: kindaEnvFileData.MaxConnLifeTime,
		MaxConnIdleTime: kindaEnvFileData.MaxConnIdleTime,
	}
}
