package handlers

import (
	"net/http"

	like "forum-project/backend/internal/repository/likes"
)

func HandelLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := GetUserId(r)
	if userID == 0 {
		HandleError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var likeData like.Like
	decode := DecodeJson(r)

	err := decode.Decode(&likeData)
	if err != nil {
		HandleError(w, "Invalid request format: "+err.Error(), http.StatusBadRequest)
		return
	}

	likeData.User_Id = userID
	message := likeData.Add()
	if message.MessageError != "" {
		HandleError(w, message.MessageError, http.StatusBadRequest)
		return
	}

	JsoneResponse(w, message.MessageSucc, http.StatusCreated)
}

func HandelDeletLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		HandleError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var like like.DeletLikes
	decode := DecodeJson(r)

	err := decode.Decode(&like)
	if err != nil {
		HandleError(w, "Failed to delete like", http.StatusBadRequest)
		return
	}

	like.DeletLike()
	JsoneResponse(w, "DELETED Like", http.StatusCreated)
}
