package db

import (
	"TueKan-backend/config"
	"database/sql"
	"fmt"

	// Postgres driver
	_ "github.com/lib/pq"
)

var db *sql.DB

// Init establish connection with database
func Init(config *config.Config) error {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s", config.DB, config.DBUser, config.DBPass, config.DBHost, config.DBPort)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return nil
}
