package handlers

import (
	"encoding/json"
	"net/http"

	user "forum-project/backend/internal/repository/users"
)

func HandleIsLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, r, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cookies, _ := r.Cookie("token")
	is := user.CheckAuthenticat(cookies.Value)
	json.NewEncoder(w).Encode(is)
}
