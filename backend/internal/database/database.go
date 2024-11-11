package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitDB() error {
	if _, err := os.Stat("../../app.db"); os.IsNotExist(err) {
		fmt.Println("Creating new database file...")
		db, err := sql.Open("sqlite3", "../../app.db")
		if err != nil {
			return err
		}
		defer db.Close()
		sqlFile, err := os.ReadFile("../internal/database/database.sql")
		if err != nil {
			return err
		}
		_, err = db.Exec(string(sqlFile))
		if err != nil {
			return err
		}
	}
	return nil
}

func Config() *sql.DB {
	db, err := sql.Open("sqlite3", "../../app.db")
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	return db
}
