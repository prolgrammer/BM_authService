package config

import "auth/config/pg"

type (
	Config struct {
		Http HTTP
		PG   pg.Config
	}
	HTTP struct {
		Host string
		Port string
	}
)

func NewConfig() *Config {
	cfg := &Config{
		Http: HTTP{
			Host: "localhost",
			Port: "8080",
		},
		PG: *pg.LoadConfig(),
	}
	return cfg
}
