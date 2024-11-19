package handlers

import (
	"fmt"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	id_user := GetUserId(r)
	fmt.Println(id_user)
	post := posts.Post{}
	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := post.Add()
	for _, name := range post.Category {
		category.AddCategory(id, name)
	}
	fmt.Println(post)
}
