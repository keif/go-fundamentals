package db

// _ tells go we need it, to prevent it from being removed
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// pointer to database type
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err) // Could not connect to database
	}

	DB.SetMaxOpenConns(10) // config connection pool
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err) // Could not create events table
	}
}
