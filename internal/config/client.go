package config

import (
	"fmt"
)

type ClientCfg struct {
	ServerHost string
	ServerPort string
}

func MustNewClientCfg() *ClientCfg {
	return &ClientCfg{
		ServerHost: MustGetEnv("SERVER_HOST"),
		ServerPort: MustGetEnv("SERVER_PORT"),
	}
}

func (c *ClientCfg) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}
