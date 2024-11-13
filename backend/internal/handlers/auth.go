package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum-project/backend/internal/models/users"
	"forum-project/backend/internal/repository"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	message := repository.Register(&user)
	if message.ErrorBool {
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
	fmt.Println(user)
	repository.Login(&user)
}
