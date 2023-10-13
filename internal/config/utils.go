package config

import (
	"fmt"
	"os"
)

func MustGetEnv(env string) string {
	v := os.Getenv(env)
	if v == "" {
		panic(fmt.Sprintf("env %s is required", env))
	}
	return v
}
