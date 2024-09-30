package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the Users table
	//     createTableSQL := `CREATE TABLE likes (
	//     id INTEGER PRIMARY KEY AUTOINCREMENT,
	//     user_id INTEGER,
	//     post_id INTEGER,
	//     comment_id INTEGER,
	//     is_like BOOLEAN NOT NULL,
	//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	//     FOREIGN KEY (user_id) REFERENCES users(id),
	//     FOREIGN KEY (post_id) REFERENCES posts(id),
	//     FOREIGN KEY (comment_id) REFERENCES comments(id),
	//     UNIQUE(user_id, post_id, comment_id)
	// );
	// `

	// Execute the SQL statement
	// statement, err := db.Prepare(createTableSQL)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// _, err = statement.Exec()
	// if err != nil {
	//     log.Fatal(err)
	// }
	insertSQL := `INSERT INTO users (username, email) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, "John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Users table created successfully!")
}
