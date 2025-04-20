package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置
type Config struct {
	// 数据库配置
	DatabaseURL string

	// 服务器配置
	ServerPort string

	// JWT配置
	JWTSecret string
}

// Load 加载配置
func Load() (*Config, error) {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("加载 .env 文件失败: %w", err)
	}

	config := &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/undb?sslmode=disable"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
	}

	return config, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
