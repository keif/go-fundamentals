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
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email VARCHAR(255) UNIQUE NOT NULL,
    	password VARCHAR(255) NOT NULL,
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create Users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER
	  	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create Events table")
	}
}
