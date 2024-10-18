package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/internal/database"
	"forum-project/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.HashPassword()
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var Responseuser struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	err = database.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", loginData.Email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !user.CheckPassword(loginData.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// Here you would typically create a session or generate a JWT token
	// For simplicity, we'll just return a success message
	Responseuser.Id = int(user.ID)
	Responseuser.Username = user.Username
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Responseuser)
}
