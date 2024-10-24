package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum-project/internal/database"
	"forum-project/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		fmt.Println("Error Parssing")
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	fmt.Println(user.Username,username)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err := user.HashPassword()
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
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
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
