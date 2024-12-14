package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"forum-project/backend/internal/database"
	"forum-project/backend/internal/handlers"
)

func main() {
	Err := database.InitDB()
	if Err != nil {
		log.Fatal(fmt.Errorf("failed to initialize database: %w", Err))
	}

	mux := http.NewServeMux()

	setupAPIRoutes(mux)
	setupPageRoutes(mux)

	serverAddr := ":3333"
	log.Printf("Server running at http://localhost%s/home\n", serverAddr)
	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func setupAPIRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/register", handlers.HandleRegister)
	mux.HandleFunc("/api/home", handlers.HomeHandle)
	mux.HandleFunc("/api/category", handlers.HandelCategory)
	mux.HandleFunc("/api/login", handlers.HandleLogin)
	mux.HandleFunc("/api/comment", handlers.Handel_GetCommet)
	mux.HandleFunc("/api/card", handlers.GetCard_handler)

	mux.Handle("/api/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.LikesHandle))))
	mux.Handle("/api/profile/posts", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfilePosts))))
	mux.Handle("/api/profile/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfileLikes))))
	mux.Handle("/api/post", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandlePost)))
	mux.Handle("/api/addcomment", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.Handler_AddComment)))
	mux.Handle("/api/like", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelLike)))
	mux.Handle("/api/deleted", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelDeletLike)))
	mux.Handle("/api/logout", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandleLogOut)))
}

func setupPageRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("../../frontend/static"))))

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		if err != nil || cookies == nil {
			http.ServeFile(w, r, "../../frontend/templates/register.html")
		} else {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		if err != nil || cookies == nil {
			http.ServeFile(w, r, "../../frontend/templates/login.html")
		} else {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/about.html")
	})

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/home.html")
	})
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/categories.html")
	})
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/contact.html")
	})
	mux.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/comment.html")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleError(w, r, http.StatusText(404), 404)
	})
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		if err != nil || cookies == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.ServeFile(w, r, "../../frontend/templates/profile.html")
		}
	})
	mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		if err != nil || cookies == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.ServeFile(w, r, "../../frontend/templates/settings.html")
		}
	})

	mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		status_code, err := strconv.Atoi(code)

		if err != nil || (status_code != 404 && status_code != 400 && status_code != 500) {
			status_code = http.StatusInternalServerError
		}
		// user := http.Cookie{
		// 	Name:  "status",
		// 	Value: strconv.Itoa(status_code),
		// }
		// http.SetCookie(w, &user)

		w.WriteHeader(status_code)
		filePath := "../../frontend/templates/err.html"
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			handlers.JsoneResponse(w, r, "Error loading the error page", http.StatusInternalServerError)
			return
		}
		// Write the file content to the ResponseWriter
		w.Write(fileContent)
	})
}
