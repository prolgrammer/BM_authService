package pg

import "time"

type Config struct {
	MaxPoolSize    int    `mapstructure:"max_pool_size"`
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"database"`
	SSLMode        string `mapstructure:"ssl_mode"`
	MigrationsPath string `mapstructure:"migrations_path"`

	RetryConnectionAttempts int
	RetryConnectionTimeout  time.Duration
}
