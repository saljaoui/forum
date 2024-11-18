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
	if r.Method != http.MethodPost {
		jsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user := repository.User{}
	err := decodeJson(r, user)
	if err != nil {
		jsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	message := user.Register()
	if message.MessageError != "" {
		jsoneResponse(w, message.MessageError, http.StatusBadRequest)
	} else {
		jsoneResponse(w, message.MessageSucc, http.StatusOK)
	}
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user := repository.Login{}
	err := decodeJson(r, user)
	
	if err != nil {
		jsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	loged, message, uuid := user.Authentication()
	if message.MessageError != "" {
		jsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return

	} else {
		SetCookie(w, "token", uuid.String(), time.Now().Add(10*time.Second))
		SetCookie(w, "user_id", fmt.Sprint(loged.Id), time.Now().Add(10*time.Second))
		jsoneResponse(w, message.MessageSucc, http.StatusOK)
	}
}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	logout := repository.Login{}
	err := decodeJson(r, logout)
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
	SetCookie(w, "token", "", time.Now())
	SetCookie(w, "user_id", "", time.Now())
}

func SetCookie(w http.ResponseWriter, name string, value string, time time.Time) {
	user := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time,
		Path:    "/",
	}
	http.SetCookie(w, &user)
}

func decodeJson(r *http.Request, user any) error {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	err := decode.Decode(&user)
	return err
}

func jsoneResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": message,
	})
}
