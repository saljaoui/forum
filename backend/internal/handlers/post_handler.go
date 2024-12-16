package handlers

import (
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, r, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_user := GetUserId(r)
	post := posts.Post{}

	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	for _, n := range post.Name_Category {
		if !checkGategory(n) {
			JsoneResponse(w, r, "Your category is incorrect", http.StatusBadRequest)
			return
		}
	}
	post.User_Id = id_user
	post.CheckPostErr(w)
	id := post.Add()

	for _, name := range post.Name_Category {
		err := category.AddCategory(id, name)
		if err != nil {
			JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
			return
		}
	}
	JsoneResponse(w, r, "create post Seccessfuly", http.StatusCreated)
}

func checkGategory(name string) bool {
	cate := []string{
		"General",
		"Technology",
		"Sports",
		"Entertainment",
		"Science",
		"Health",
		"Food",
		"Travel",
		"Fashion",
		"Art",
		"Music",
	}
	for _, v := range cate {
		if v == name {
			return true
		}
	}
	return false
}
