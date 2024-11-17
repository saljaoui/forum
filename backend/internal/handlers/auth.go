package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	repository "forum-project/backend/internal/repository/users"
)

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
		json.NewEncoder(w).Encode(messages.MessageSucc)
		next.ServeHTTP(w, r)
	})
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	user := repository.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	message := user.Register()
	if message.MessageError != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(string(message.MessageError))
	} else {
		json.NewEncoder(w).Encode(string(message.MessageSucc))
	}
	w.Header().Set("Content-Type", "application/json")
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	user := repository.Login{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error to login")
		return
	}
	loged, message, uuid := user.Authentication()
	if message.MessageError != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.MessageError)
		return

	} else {
		user := http.Cookie{
			Name:    "token",
			Value:   uuid.String(),
			Expires: time.Now().Add(10 * time.Second),
			Path:    "/",
		}
		http.SetCookie(w, &user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loged)
	}
}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	id := r.URL.RawQuery
	fmt.Print(id)
	// logout := repository.Login{}
	// iduser, err := strconv.Atoi(id)
	// if err != nil {
	// 	fmt.Println("error to get id user")
	// }
	// logout.Id = int64(iduser)
	// logout.LogOut()
}

func DisplyPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcom to Page Home")
	fmt.Fprintln(w, "welcom")
}
