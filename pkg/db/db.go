package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() *sql.DB {
	//connStr := "your_connection_string_here"
	db, err := sql.Open("sqlite3", "./data/main.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Database connected.")
	return db
}
