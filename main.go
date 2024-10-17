package main

import (
	"log"
	"net/http"

	"forum-project/internal/database"
	"forum-project/internal/handlers"

)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}


	// API routes
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/login", handlers.LoginHandler)

	// Serve static files
	

	// Serve HTML templates
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/register.html")
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
