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
	http.HandleFunc("/api/home", handlers.Dachboard)
	http.HandleFunc("/api/createPost", handlers.CreateNewPost)
	http.HandleFunc("/api/getPosts", handlers.GetPostsHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Serve HTML templates
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/register.html")
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/home.html")
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
