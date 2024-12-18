package config

import (
	"os"
)

type Config struct {
	Addr string
}

func GettingConfig() *Config {
	config := new(Config)
	address := os.Getenv("PORT")
	config.Addr = address
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}
