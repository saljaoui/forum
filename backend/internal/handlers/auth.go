package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum-project/backend/internal/models"
	"forum-project/backend/internal/repository"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := models.User{}
	db := repository.Connect{}  

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error decoding JSON:")
		//http.Error(w, `{"error": "Invalid JSON data"}`, http.StatusBadRequest)
		return
	}
	err = db.Register(&user)
	if err != nil {
		fmt.Println("register:")
	}
}
