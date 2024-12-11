package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	like "forum-project/backend/internal/repository/likes"
)

func HandelLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	id_user := GetUserId(r)
	like := like.Like{}
	decode := DecodeJson(r)
	err := decode.Decode(&like)
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	like.User_Id = id_user
	m := like.Add()
	if m.MessageError != "" {
		HandleError(w, m.MessageError, http.StatusBadRequest)
		return
	}
	JsoneResponse(w, m.MessageSucc, http.StatusCreated)
}

func HandelDeletLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
 	like := like.DeletLikes{}
	decode := DecodeJson(r)
	err := decode.Decode(&like)
	if err != nil {
		fmt.Println(err)
		HandleError(w, "err.Error()", http.StatusBadRequest)
		return
	}
  
 	like.DeletLike()
 	JsoneResponse(w, "DELETED Like", http.StatusCreated)
}
