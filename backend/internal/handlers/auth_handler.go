package handlers

import (
	"fmt"
	"net/http"
	"time"

	repository "forum-project/backend/internal/repository/users"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user repository.User
	decode := DecodeJson(r)
	decode.DisallowUnknownFields()

	err := decode.Decode(&user)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRegiseter, message, uuid := user.Register()
	if message.MessageError != "" {
		HandleError(w, message.MessageError, http.StatusBadRequest)
		return
	}

	SetCookie(w, "token", uuid, time.Now().Add(2*time.Minute))
	// SetCookie(w, "user_id", fmt.Sprint(userRegiseter.Id), time.Now().Add(2*time.Minute))
	JsoneResponse(w, userRegiseter, http.StatusOK)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user repository.Login
	decode := DecodeJson(r)
	err := decode.Decode(&user)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}
	loged, message, uuid := user.Authentication()
	user.Getuuid(uuid.String())
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
		// HandleError(w, message.MessageError, http.StatusBadRequest)
		return
	}

	SetCookie(w, "token", uuid.String(), time.Now().Add(1*time.Hour))
	// SetCookie(w, "user_id", fmt.Sprint(loged.Id), time.Now().Add(1*time.Hour))
	JsoneResponse(w, loged, http.StatusOK)
}

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var logout repository.Logout
	decode := DecodeJson(r)

	err := decode.Decode(&logout)
	if err != nil {
		HandleError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	logout.Id = int64(GetUserId(r))
	var uuid repository.UUID

	message := uuid.UUiduser(logout.Uuid)
	if message.MessageError != "" {
		HandleError(w, "Missing or invalid Uuid", http.StatusBadRequest)
		return
	}

	message = logout.LogOut()
	if message.MessageError != "" {
		HandleError(w, message.MessageError, http.StatusBadRequest)
		return
	}

	clearCookies(w)
	w.WriteHeader(http.StatusOK)
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

func GetUserId(r *http.Request) int {
	cookie, err := r.Cookie("token")
	if err != nil {
		return 0
	}
	uuid := repository.UUID{}
	m := uuid.UUiduser(cookie.Value)
	if m.MessageError != "" {
		fmt.Println(m.MessageError)
	}

	return uuid.Iduser
}

func clearCookies(w http.ResponseWriter) {
	SetCookie(w, "token", "", time.Now())
	SetCookie(w, "user_id", "", time.Now())
}
