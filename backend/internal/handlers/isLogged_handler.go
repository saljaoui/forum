package handlers

import (
	"net/http"

	repository "forum-project/backend/internal/repository/users"
)

func HandleIsLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, r, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user := repository.User{}
	cookies, _ := r.Cookie("token")
	is := user.AuthenticatLogin(cookies.Value)
	if is.MessageError != "" {
		
	}
	JsoneResponse(w, r, is, http.StatusOK)
	// json.NewEncoder(w).Encode(isLogged.IsItLogged)
}
