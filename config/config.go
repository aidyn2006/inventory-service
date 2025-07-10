package config

import "os"

type Config struct {
	DB         DBConfig
	ServerPort string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() *Config {
	return &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "${PASSWORD}"),
			DBName:   getEnv("DB_NAME", "go"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		ServerPort: getEnv("SERVER_PORT", "8082"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
