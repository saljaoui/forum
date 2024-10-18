
package handlers

import (
	"net/http"
	"encoding/json"
    "time"
    "forum-project/internal/models"
    "forum-project/internal/database"
    "fmt"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request)  {
	var createPost models.CreatePost
    err := json.NewDecoder(r.Body).Decode(&createPost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if createPost.Title == "" || createPost.Content == "" {
        http.Error(w, "Title and content are required", http.StatusBadRequest)
        return
    }

	_, err = database.DB.Exec("INSERT INTO posts (user_id, title, content, created_at) VALUES (?, ?, ?, ?)",
	createPost.UserID, createPost.Title, createPost.Content, time.Now())
if err != nil {
	http.Error(w, "Error creating post", http.StatusInternalServerError)
	return
}
w.WriteHeader(http.StatusCreated)
// json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})

}