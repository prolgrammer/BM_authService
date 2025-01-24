package config

import (
	"auth/config/pg"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		App         `mapstructure:"app"`
		TokenConfig `mapstructure:"token_config"`
		HTTP        `mapstructure:"http"`
		JWT         `mapstructure:"jwt"`
		PG          pg.Config `mapstructure:"postgres"`
	}

	App struct {
		GinMode string `mapstructure:"gin_mode"`
	}

	TokenConfig struct {
		AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
		RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
	}

	HTTP struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	JWT struct {
		SignSecretToken string `mapstructure:"secret_key"`
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: could not load .env file: %s\n", err)
	}

	cfg := Config{}
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config/")

	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	for _, key := range v.AllKeys() {
		anyValue := v.Get(key)
		str, ok := anyValue.(string)
		if !ok {
			continue
		}

		replaced := os.ExpandEnv(str)
		v.Set(key, replaced)
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling file: %s", err))
	}

	return &cfg, nil
}
