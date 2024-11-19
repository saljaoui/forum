package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	post := posts.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	post.CheckPostErr(w)
	id := post.Add()
	fmt.Println(post.Name_Category)
	for _,name := range post.Name_Category {
		category.AddCategory(id, name)
	}
	fmt.Println(post)
}
