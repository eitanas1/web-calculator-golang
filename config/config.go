package config

import (
	"os"
)

type Config struct {
	Host string
	Port string
}

func GettingConfig() *Config {
	config := new(Config)

	config.Host = os.Getenv("HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}

	return config
}
