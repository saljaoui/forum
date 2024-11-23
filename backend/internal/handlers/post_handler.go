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
	id_user := GetUserId(r)
	fmt.Println(id_user)
	post := posts.Post{}
	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.User_Id = id_user
	post.CheckPostErr(w)
	id := post.Add()
	for _, name := range post.Name_Category {
		category.AddCategory(id, name)
	}
}
