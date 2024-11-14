package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum-project/backend/internal/models"
	repository "forum-project/backend/internal/repository/users"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		if err != nil || cookies == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		sessionID := cookies.Value
		if sessionID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	message := repository.Register(&user)
	if message.MessageError != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(string(message.MessageError))
	} else {
		json.NewEncoder(w).Encode(string(message.MessageSucc))
	}
	w.Header().Set("Content-Type", "application/json")
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	user := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error to login")
		return
	}

	loged, message, uuid := repository.Login(&user)

	if message.MessageError != "" {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.MessageError)

		return

	} else {
		user := http.Cookie{
			Name:    "token",
			Value:   uuid.String(),
			Expires: time.Now().Add(5 * time.Second),
		}
		http.SetCookie(w, &user)
		fmt.Println(user.Value)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loged)
	}
}
