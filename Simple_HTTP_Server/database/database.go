package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./items.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table if it doesn't exist
	createTable := `
    CREATE TABLE IF NOT EXISTS items (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL
    );`

	_, err = DB.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}
}
