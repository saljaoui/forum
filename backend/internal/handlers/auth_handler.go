package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	repository "forum-project/backend/internal/repository/users"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user := repository.User{}
	decode := DecodeJson(r)
	err := decode.Decode(&user)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	message := user.Register()
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
	} else {
		JsoneResponse(w, message.MessageSucc, http.StatusCreated)
	}
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user := repository.Login{}
	decode := DecodeJson(r)
	err := decode.Decode(&user)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user.Email)
	loged, message, uuid := user.Authentication()
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return
	} else {
		SetCookie(w, "token", uuid.String(), time.Now().Add(10*time.Second))
		SetCookie(w, "user_id", fmt.Sprint(loged.Id), time.Now().Add(10*time.Second))
		JsoneResponse(w, "User Login successfully", http.StatusOK)
	}
}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	logout := repository.Login{}
	decode := DecodeJson(r)
	err := decode.Decode(&logout)
	if err != nil {
		JsoneResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	jsonValue, err := r.Cookie("user_id")
	if err != nil {
		JsoneResponse(w, "Missing or invalid user_id cookie", http.StatusBadRequest)
		return
	}
	user_id, err := strconv.Atoi(jsonValue.Value)
	if err != nil {
		JsoneResponse(w, "Invalid user_id value", http.StatusBadRequest)
		return
	}
	if int64(user_id) != logout.Id {
		JsoneResponse(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	message := logout.LogOut()
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
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
