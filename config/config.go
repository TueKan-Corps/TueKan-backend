package config

import (
	"os"
)

// Config basic config of an app
type Config struct {
	DBHost       string
	DBUser       string
	DBPort       string
	DBPass       string
	DB           string
	Port         string
	AmazonID     string
	AmazonSecret string
}

func (config *Config) Init() error {

	config.DBHost = os.Getenv("DB_HOST")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPass = os.Getenv("DB_PASS")
	config.DB = os.Getenv("DB")
	config.AmazonID = os.Getenv("AWS_ACCESS_ID")
	config.AmazonSecret = os.Getenv("AWS_ACCESS_SECRET")

	var ok bool
	config.Port, ok = os.LookupEnv("PORT")

	if ok == false {
		config.Port = "1323"
	}

	if _, err := os.Stat("./img"); os.IsNotExist(err) {
		os.Mkdir("./img", 0777)
	}

	return nil
}
