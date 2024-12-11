package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/backend/internal/repository/profile"
)

func HandleProfilePosts(w http.ResponseWriter, r *http.Request) {
	id_user := GetUserId(r)
	posts := profile.GetPostsProfile(id_user)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		HandleError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func HandleProfileLikes(w http.ResponseWriter, r *http.Request) {
	id_user := GetUserId(r)
	posts := profile.GetPostsProfileByLikes(id_user)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		HandleError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
