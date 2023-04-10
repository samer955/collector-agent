package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DatabaseConfig struct {
	Connection *sql.DB
}

var dbConfig DatabaseConfig

func GetDbConfig() DatabaseConfig {
	return dbConfig
}

func init() {

	if err := godotenv.Load("config.env"); err != nil {
		log.Println("Error loading .env file.", err)
	}

	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := sql.Open(DB_DRIVER, dsn)
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println("Unable to connect to DB. Metrics will NOT be stored. ", err)
		return
	}

	dbConfig.Connection = db

	log.Println("connected to Database")

}
