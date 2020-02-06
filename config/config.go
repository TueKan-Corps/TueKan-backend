package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config basic config of an app
type Config struct {
	DBHost string
	DBUser string
	DBPort string
	DBPass string
	DB     string
}

func (config *Config) Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	config.DBHost = os.Getenv("DB_HOST")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPass = os.Getenv("DB_PASS")
	config.DB = os.Getenv("DB")

	return nil
}
