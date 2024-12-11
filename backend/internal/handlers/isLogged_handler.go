package handlers

import (
	"net/http"

	user "forum-project/backend/internal/repository/users"
)

func HandleIsLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cookies, _ := r.Cookie("token")
	is, err := user.IsLogged(cookies.Value)
	if err != nil {
		HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsoneResponse(w, is, http.StatusOK)
	// json.NewEncoder(w).Encode(isLogged.IsItLogged)
}
