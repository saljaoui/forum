package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	repository "forum-project/backend/internal/repository/users"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Middleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET , POST , OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	response := Response{
		Message: "Hello, World!",
		Status:  "success",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		user := repository.User{}
		if err != nil || cookies == nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized: Cookie not present", http.StatusUnauthorized)
				fmt.Println("Unauthorized: Cookie not present")
				return
			}
		}
		if cookies.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		messages := user.AuthenticatLogin(cookies.Value)
		if messages.MessageError != "" {
			json.NewEncoder(w).Encode(messages.MessageError)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
