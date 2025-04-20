package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	JWT        JWTConfig        `yaml:"jwt"`
	API        APIConfig        `yaml:"api"` // Added API configuration
	Security   SecurityConfig   `yaml:"security"` // Added Security configuration

}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// APIConfig API配置
type APIConfig struct {
	EnableRateLimit bool          `yaml:"enableRateLimit"`
	RateLimit      int           `yaml:"rateLimit"`
	RateWindow     time.Duration `yaml:"rateWindow"`
	MaxPageSize    int           `yaml:"maxPageSize"`
	DefaultVersion string        `yaml:"defaultVersion"`
}

// SecurityConfig Security配置
type SecurityConfig struct {
	JWTSecret      string        `yaml:"jwtSecret"`
	TokenExpiry    time.Duration `yaml:"tokenExpiry"`
	AllowedOrigins []string      `yaml:"allowedOrigins"`
}


// Load 加载配置文件
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}