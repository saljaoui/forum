package handlers

import (
	"encoding/json"
	"fmt"
	"forum-project/backend/internal/repository/posts"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	post := posts.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	post.Add()
	fmt.Println(post)
}
