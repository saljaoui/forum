package main

import (
	"fmt"
	"log"
	"net/http"

	"forum-project/backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.TestHandlers)
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/static"))))    
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/login.html")
		// fmt.Println("ok")
	})

    fmt.Println("Server running at :3333")
    fmt.Println("http://localhost:3333/")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
