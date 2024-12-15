package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	user "forum-project/backend/internal/repository/users"
)

func HandleIsLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, r, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cookies, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}
 	is := user.CheckAuthenticat(cookies.Value)
 	json.NewEncoder(w).Encode(is)
}
