package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewDB() {
	var err error
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)

	DB, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	log.Println("Successfully connected to database")
}