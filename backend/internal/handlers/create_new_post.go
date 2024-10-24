package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/internal/database"
	"forum-project/internal/models"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("INSERT INTO posts (user_id, title, content)VALUES(?,? ,?)", post.User_id, post.Title, post.Content)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	_, err = database.DB.Exec("INSERT INTO categories (name)VALUES(?)", post.Category)
	if err != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
