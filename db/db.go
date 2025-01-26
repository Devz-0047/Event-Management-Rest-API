package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Initialize the global DB variable properly
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	// Set connection pooling limits
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create tables
	err = createTables()
	if err != nil {
		panic(fmt.Sprintf("Could not create tables: %v", err))
	}
}

func createTables() error {
	// Define the SQL statement to create the events table
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER
	);
	`

	// Execute the SQL statement
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		return fmt.Errorf("error creating events table: %w", err)
	}

	return nil
}
