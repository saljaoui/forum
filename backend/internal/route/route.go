package route

import (
	"net/http"
	"os"
	"strings"

	"forum-project/backend/internal/handlers"
)

func SetupAPIRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/register", handlers.HandleRegister)
	mux.HandleFunc("/api/login", handlers.HandleLogin)
	mux.HandleFunc("/api/home", handlers.HomeHandle)
	mux.HandleFunc("/api/category", handlers.HandelCategory)
	mux.HandleFunc("/api/comment", handlers.Handel_GetCommet)
	mux.HandleFunc("/api/card", handlers.GetCard_handler)
	mux.HandleFunc("/api/isLogged", handlers.HandleIsLogged)
	mux.Handle("/api/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.LikesHandle))))
	mux.Handle("/api/profile/posts", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfilePosts))))
	mux.Handle("/api/profile/likes", handlers.AuthenticateMiddleware((http.HandlerFunc(handlers.HandleProfileLikes))))
	mux.Handle("/api/post", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandlePost)))
	mux.Handle("/api/addcomment", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.Handler_AddComment)))
	mux.Handle("/api/like", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelLike)))
	mux.Handle("/api/deleted", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandelDeletLike)))
	mux.Handle("/api/logout", handlers.AuthenticateMiddleware(http.HandlerFunc(handlers.HandleLogOut)))
}

func SetupPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
 		if r.Method != http.MethodGet {
			handlers.JsoneResponseError(w, r, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

 		suffix := r.URL.Path[len("/static/"):]

 		if strings.Contains(suffix, ".css/") || strings.Contains(suffix, ".js/") || strings.Contains(suffix, ".png/") {
			handlers.JsoneResponseError(w, r, "Not Found", http.StatusNotFound)
			return
		}
 
		if strings.Contains(suffix, ".js") {
			http.ServeFile(w, r, "../../frontend/static/"+suffix)
			return
		}

 		allowedFiles := map[string]bool{
			"css/alert.css":       true,
			"css/styles.css":      true,
			"imgs/logo.png":       true,
			"imgs/profilePic.png": true,
		}

 		if !allowedFiles[suffix] {
			handlers.JsoneResponseError(w, r, "Access Forbidden", http.StatusForbidden)
			return
		} 
		http.ServeFile(w, r, "../../frontend/static/"+suffix)
	})

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
		validatePath(w, r)
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
		filePath := "../../frontend/templates/err.html"
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			handlers.JsoneResponse(w, r, "Error loading the error page", http.StatusInternalServerError)
			return
		}
		w.Write(fileContent)
	})
}

func isValidPath(path string, paths []string) bool {
	for _, v := range paths {
		if path == v {
			return true
		}
	}
	return false
}

// this is validate path function
func validatePath(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"/comment",
		"/register",
		"/login",
		"/logout",
		"/about",
		"/contact",
		"/home",
		"/categories",
		"/profile",
		"/settings",
		"/err",
	}
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	} else if !isValidPath(r.URL.Path, paths) || r.URL.Path == "/logout" {
		handlers.JsoneResponseError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}
}
