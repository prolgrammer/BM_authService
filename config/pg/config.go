package pg

import "time"

type Config struct {
	MaxPoolSize    int
	Host           string
	Port           int
	User           string //postgres
	Password       string //1234
	Database       string
	SSLMode        string
	MigrationsPath string

	RetryConnectionAttempts int
	RetryConnectionTimeout  time.Duration
}

func LoadConfig() *Config {
	cfg := &Config{
		Host:           "localhost",
		Port:           5432,
		User:           "postgres",
		Password:       "1234",
		Database:       "BoxMaster",
		SSLMode:        "disable",
		MigrationsPath: "file://config/pg/migrations",
	}
	return cfg
}
