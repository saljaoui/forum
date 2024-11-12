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

func SelectOneRow(query string, model any, st ...any) {
	db := Config()
	err := db.QueryRow(query, model).Scan(st...)
	if err != nil {
		fmt.Println(err)
	}
}

func SelectRows(query string, model ...any) *sql.Rows {
	db := Config()
	rows, err := db.Query(query, model...)
	if err != nil {
		fmt.Println(err)
	}

	return rows
}

func Exec(query string, model ...any) {
	db := Config()
	_, err := db.Exec(query, model...)
	if err != nil {
		fmt.Println("error here:", err)
	}
}
