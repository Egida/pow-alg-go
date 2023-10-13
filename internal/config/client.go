package config

type ClientCfg struct {
	ServerAddr string
}

func MustNewClientCfg() *ClientCfg {
	return &ClientCfg{
		ServerAddr: MustGetEnv("SERVER_ADDR"),
	}
}
