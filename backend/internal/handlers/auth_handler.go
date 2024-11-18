package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	repository "forum-project/backend/internal/repository/users"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	user := repository.User{}
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	err := decode.Decode(&user)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}
	message := user.Register()
	if message.MessageError != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(string(message.MessageError))
	} else {
		json.NewEncoder(w).Encode(string(message.MessageSucc))
	}
	w.Header().Set("Content-Type", "application/json")
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	user := repository.Login{}
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	err := decode.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}
	loged, message, uuid := user.Authentication()
	if message.MessageError != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.MessageError)
		return

	} else {

		user := http.Cookie{
			Name:    "token",
			Value:   uuid.String(),
			Expires: time.Now().Add(30 * time.Second),
			Path:    "/",
		}

		user_id := http.Cookie{
			Name:    "user_id",
			Value:   fmt.Sprint(loged.Id),
			Expires: time.Now().Add(30 * time.Second),
			Path:    "/",
		}
		http.SetCookie(w, &user)
		http.SetCookie(w, &user_id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loged)
	}
}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	logout := repository.Login{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&logout)
	if err != nil {
		jsoneResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	jsonValue, err := r.Cookie("user_id")
	if err != nil {
		jsoneResponse(w, "Missing or invalid user_id cookie", http.StatusBadRequest)
		return
 	}
	user_id, err := strconv.Atoi(jsonValue.Value)
	if err != nil {
		jsoneResponse(w, "Invalid user_id value", http.StatusBadRequest)
		return
 	}
	if int64(user_id) != logout.Id {
		jsoneResponse(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	message := logout.LogOut()
	if message.MessageError != "" {
		jsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return
	}
	cookieLogOut(w)
}

func cookieLogOut(w http.ResponseWriter) {
	user := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
	}
	http.SetCookie(w, &user)
}

func jsoneResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": message,
	})
}
