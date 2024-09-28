package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Open a connection to the SQLite database
    db, err := sql.Open("sqlite3", "./databas.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the Users table
    createTableSQL := `CREATE TABLE IF NOT EXISTS Users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`

    // Execute the SQL statement
    statement, err := db.Prepare(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    _, err = statement.Exec()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Users table created successfully!")
}
