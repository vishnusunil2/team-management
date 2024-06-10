package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerAddress string
	DbDSN         string
	LogLevel      string
}

var cfg *Config

func LoadConfig() (*Config, error) {
	//Load environment variables from .env in root folder
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}
	//Create and populate config struct
	cfg = &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DbDSN:         os.Getenv("DB_KEY"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
	}
	return cfg, nil
}
