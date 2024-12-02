package handlers

import (
	"encoding/json"
	"net/http"

	category "forum-project/backend/internal/repository/categories"

)

func HandelCategory(w http.ResponseWriter, r *http.Request) {
	categoryStruct := category.Category{}
	decode := DecodeJson(r)
	err := decode.Decode(&categoryStruct)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts := category.GetPostsByCategoryId(categoryStruct.Id)
	json.NewEncoder(w).Encode(posts)
}
