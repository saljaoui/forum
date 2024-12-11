package handlers

import (
	"net/http"

	repository "forum-project/backend/internal/repository/users"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		var user repository.User
		if err != nil || cookies == nil {
			if err == http.ErrNoCookie {
				HandleError(w, "Unauthorized: Cookie not presen", http.StatusUnauthorized)
				return
			}
		}

		if cookies.Value == "" {
			HandleError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		messages := user.AuthenticatLogin(cookies.Value)
		if messages.MessageError != "" {
			HandleError(w, messages.MessageError, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
