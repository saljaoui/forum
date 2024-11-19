package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	post := posts.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	id := post.Add()
	fmt.Println(post.Category)
	for _, name := range post.Category {
		category.AddCategory(id, name)
	}
	fmt.Println(post)
}
