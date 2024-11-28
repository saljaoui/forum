package handlers

import (
	"encoding/json"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/home"
)

func HandelCategory(w http.ResponseWriter, r *http.Request) {
	category := category.Category{}
	decode := DecodeJson(r)
	err := decode.Decode(&category)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts := home.GetPosts(category.Id)
	json.NewEncoder(w).Encode(posts)
}
