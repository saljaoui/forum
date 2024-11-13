package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

<<<<<<< HEAD
	models "forum-project/backend/internal/models/posts"
	"forum-project/backend/internal/repository"
=======
	"forum-project/backend/internal/models"
	"forum-project/backend/internal/repository/posts"
>>>>>>> soufian
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON Post:", err)
		return
	}
	// fmt.Println(post)
	message := posts.Post(&post)
	if message.ErrorBool {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(string(message.MessageError))
	} else {
		json.NewEncoder(w).Encode(string(message.MessageSucc))
	}
	w.Header().Set("Content-Type", "application/json")
}
