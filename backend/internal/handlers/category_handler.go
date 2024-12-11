package handlers

import (
	"encoding/json"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
)

func HandelCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var categoryStruct category.Category
	decode := DecodeJson(r)

	err := decode.Decode(&categoryStruct)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts := category.GetPostsByCategoryId(categoryStruct.Category)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		HandleError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
