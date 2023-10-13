package config

type ServerCfg struct {
	Addr string
}

func MustNewServerCfg() *ServerCfg {
	return &ServerCfg{
		Addr: MustGetEnv("SERVER_ADDR"),
	}
}
