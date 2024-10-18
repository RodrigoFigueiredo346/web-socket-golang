package db

import (
	"database/sql"
	"fmt"
	"log"
	"main/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	cfg := config.GetConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.DbPort, cfg.User, cfg.Password, cfg.Dbname)

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Database connection established successfully!")
}
