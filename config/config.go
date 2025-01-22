package config

import (
	"auth/config/pg"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Http HTTP      `mapstructure:"http"`
		PG   pg.Config `mapstructure:"postgres"`
	}
	HTTP struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
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
