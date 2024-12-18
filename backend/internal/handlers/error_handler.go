package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var errRes errsResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&errRes)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if errRes.Code == http.StatusNotFound {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
	} else if errRes.Code == http.StatusBadRequest {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
	} else if errRes.Code == http.StatusMethodNotAllowed {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
	} else if errRes.Code == http.StatusInternalServerError {
		JsoneResponse(w, r, errRes.Msg, http.StatusInternalServerError)
	}
}
