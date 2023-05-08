package config

import (
	"os"
)

type Config struct {
	ServerPort     string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

const (
	defaultServerPort = "8080"
	defaultDBName     = "postgres"
	defaultDBUser     = "postgres"
	defaultDBPW       = "postgres"
	defaultDBHost     = "localhost"
	defaultDBPort     = "5432"
)

func New() Config {
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", defaultDBHost),
		Port:     getEnv("DB_PORT", defaultDBPort),
		Name:     getEnv("DB_NAME", defaultDBName),
		User:     getEnv("DB_USER", defaultDBUser),
		Password: getEnv("DB_PASSWORD", defaultDBPW),
	}

	cfg := Config{ServerPort: defaultServerPort, DatabaseConfig: dbConfig}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
