package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/backend/internal/repository/cards"
	like "forum-project/backend/internal/repository/likes"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	posts := cards.GetAllCards()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		HandleError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func LikesHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	var liked like.Like
	decode := DecodeJson(r)
	
	w.Header().Set("Content-Type", "application/json")
	err := decode.Decode(&liked)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	dislike := liked.ChecklikesUser()
	err = json.NewEncoder(w).Encode(dislike)
	if err != nil {
		HandleError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
