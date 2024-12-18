package handlers

import (
	"fmt"
	"net/http"
	"time"

	repository "forum-project/backend/internal/repository/users"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		user := repository.User{}
		if err != nil || cookies == nil {
			if err == http.ErrNoCookie {
				JsoneResponse(w, r, "Unauthorized: Cookie not presen", http.StatusUnauthorized)
				return
			}
		}
		if cookies.Value == "" {
			// logout := repository.Logout{}
			// u := repository.UUID{}
			// u.UUiduser(cookies.Value)
			// logout.Id = int64(u.Iduser)
			// logout.Uuid = cookies.Value
			// logout.LogOut()
			// 		timeex := time.Now().Add(0* time.Second).UTC()
			// err := updateUUIDUser("null", us.Id, timeex)
			// if err != nil {
			// 	m.MessageError = "Error To Update user"
			// 	return
			// }
			JsoneResponse(w, r, "Unauthorized", http.StatusUnauthorized)
			return
		}
		messages, expire := user.AuthenticatLogin(cookies.Value)
		if messages.MessageError != "" {
			JsoneResponse(w, r, messages.MessageError, http.StatusUnauthorized)
			return
		}
		if time.Now().Before(expire) {
			fmt.Println("The current time is before the expiration time.")
		} else {
			logout := repository.Logout{}
			u := repository.UUID{}
			u.UUiduser(cookies.Value)
			logout.Id = int64(u.Iduser)
			logout.Uuid = cookies.Value
			logout.LogOut()
			  fmt.Println("The expiration time has passed. Meddel")
		}
		next.ServeHTTP(w, r)
	})
}
