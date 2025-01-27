package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err := createTables(); err != nil {
		panic(fmt.Sprintf("Could not create tables: %v", err))
	}
}

func createTables() error {
	// Create Users Table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	if _, err := DB.Exec(usersTable); err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}

	// Create Events Table
	eventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	if _, err := DB.Exec(eventsTable); err != nil {
		return fmt.Errorf("error creating events table: %w", err)
	}

	// Create Registrations Table
	registrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	if _, err := DB.Exec(registrationsTable); err != nil {
		return fmt.Errorf("error creating registrations table: %w", err)
	}

	return nil
}
