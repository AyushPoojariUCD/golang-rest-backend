package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	createTables()
}

func createTables() {

	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL
	);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal("Could not create events table:", err)
	}
}