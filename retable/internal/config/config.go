package config

type Config struct {
	PostgresURL string
	RedisURL    string
}

func LoadConfig() *Config {
	// TODO: Implement configuration loading from environment variables
	return &Config{
		PostgresURL: "postgres://postgres:postgres@localhost:5432/undb?sslmode=disable",
		RedisURL:    "redis://localhost:6379/0",
	}
}
