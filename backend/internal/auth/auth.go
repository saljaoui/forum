package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum-project/backend/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var User models.User
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&User)
	
}
