package config

import (
	"fmt"
)

type ServerCfg struct {
	Host string
	Port string
}

func MustNewServerCfg() *ServerCfg {
	return &ServerCfg{
		Host: MustGetEnv("SERVER_HOST"),
		Port: MustGetEnv("SERVER_PORT"),
	}
}

func (c *ServerCfg) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
