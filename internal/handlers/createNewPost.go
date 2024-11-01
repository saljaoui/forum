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
	_, err = database.DB.Exec("INSERT INTO posts (user_id, title, content)VALUES(?,? ,?)", post.User_id, post.Title, post.Content)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	post.ID, err = result.LastInsertId()
	if err != nil {
		http.Error(w, "Error retrieving post ID", http.StatusInternalServerError)
		return
	}
	idc, _ := strconv.Atoi(post.Category)
	_, err = database.DB.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", post.ID, idc)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, user_id, title, content, created_at FROM posts")
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	
	var posts []models.Post


	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.User_id, &post.Title, &post.Content, &post.Created_at)
		if err != nil {
			http.Error(w, "Error scanning post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
