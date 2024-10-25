package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
	"forum-project/internal/database"
	"forum-project/internal/models"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	result, err := database.DB.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", post.User_id, post.Title, post.Content)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	post.ID, err = result.LastInsertId()
	if err != nil {
		http.Error(w, "Error retrieving post ID", http.StatusInternalServerError)
		return
	}

	err = database.DB.Exec("INSERT INTO posts (post_id, title) VALUES (?, ?)", post.ID, strconv.Atoi(post.Category))
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	fmt.Println(post.ID)
	// Responseuser.Id = int(user.ID)
	// Responseuser.Username = user.Username
	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Responseuser)

	// _, err = database.DB.Exec("INSERT INTO categories (name)VALUES(?)", post.Category)
	// if err != nil {
	// 	http.Error(w, "Error creating category", http.StatusInternalServerError)
	// 	return
	// }
	w.WriteHeader(http.StatusCreated)
}
