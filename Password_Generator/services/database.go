package services

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitializeDB sets up the SQLite database and creates the table if it doesn't exist
func InitializeDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return db
}

// SavePassword stores a new password in the database
func SavePassword(db *sql.DB, password string) {
	insertSQL := `INSERT INTO passwords (password) VALUES (?)`
	_, err := db.Exec(insertSQL, password)
	if err != nil {
		log.Fatalf("Error inserting password: %v", err)
	}
}

// GetRecentPasswords retrieves the most recent 10 passwords from the database
func GetRecentPasswords(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT password FROM passwords ORDER BY created_at DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []string
	for rows.Next() {
		var password string
		if err := rows.Scan(&password); err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}
	return passwords, nil
}
