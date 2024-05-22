package db

// _ tells go we need it, to prevent it from being removed
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// pointer to database type
var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "./api.db")

	if err != nil {
		panic(err) // Could not connect to database
	}

	DB.SetMaxOpenConns(10) // config connection pool
	DB.SetMaxIdleConns(5)
}
