package handlers

import (
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		JsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_user := GetUserId(r)
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
	JsoneResponse(w, "create post Seccessfuly", http.StatusCreated)
}
