package config

import (
	"fmt"
)

type ServerCfg struct {
	Host string
	Port string

	RedisHost string
	RedisPort string
}

func MustNewServerCfg() *ServerCfg {
	return &ServerCfg{
		Host: MustGetEnv("SERVER_HOST"),
		Port: MustGetEnv("SERVER_PORT"),

		RedisHost: MustGetEnv("REDIS_HOST"),
		RedisPort: MustGetEnv("REDIS_PORT"),
	}
}

func (c *ServerCfg) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *ServerCfg) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}
