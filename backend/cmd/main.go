package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"forum-project/backend/internal/auth"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	Err := database.InitDB()
	if Err != nil {
		log.Fatal(Err)
	}

	mux.HandleFunc("/api/register", auth.RegisterHandler)

	mux.HandleFunc("/", handlers.TestHandlers)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/static"))))

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/login.html")
	})

	fmt.Println("Server running at :3333")
	fmt.Println("http://localhost:3333")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
