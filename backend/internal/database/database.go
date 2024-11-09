package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func InitDB() error {
	if _, err := os.Stat("../internal/database/app.db"); os.IsNotExist(err) {
		fmt.Println("Creating new database file...")
		DB, err := sql.Open("sqlite3", "../internal/database/app.db")
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}
		defer DB.Close()
		sqlFile, err := os.ReadFile("../internal/database/init-db.sql")
		if err != nil {
			return fmt.Errorf("failed to read SQL file: %w", err)
		}
		_, err = DB.Exec(string(sqlFile))
		if err != nil {
			return fmt.Errorf("failed to execute SQL file: %w", err)
		}
		fmt.Println("Database initialized successfully.")
	}
	return nil
}
