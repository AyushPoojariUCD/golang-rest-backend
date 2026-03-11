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
		log.Fatal("Could not open database:", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Could not create users table:", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create events table:", err)
	}
}