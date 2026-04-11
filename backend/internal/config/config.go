package config

import (
	"os"
	"strings"
)

const defaultPort = "8080"

type Config struct {
	Port string
}

func Load() Config {
	port := strings.TrimSpace(os.Getenv("PORT"))
	port = strings.TrimPrefix(port, ":")
	if port == "" {
		port = defaultPort
	}

	return Config{
		Port: port,
	}
}

func (c Config) Addr() string {
	return ":" + c.Port
}
