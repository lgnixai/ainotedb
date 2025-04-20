package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	JWTSecret   string
	LogLevel    string
}

func Load() (*Config, error) {
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/undb?sslmode=disable")
	viper.SetDefault("SERVER_PORT", "5000")
	viper.SetDefault("JWT_SECRET", "your-secret-key")
	viper.SetDefault("LOG_LEVEL", "info")

	viper.AutomaticEnv()

	config := &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		ServerPort:  viper.GetString("SERVER_PORT"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		LogLevel:    viper.GetString("LOG_LEVEL"),
	}

	return config, nil
}
