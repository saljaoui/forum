package handlers

import (
	"encoding/json"
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
		JsoneResponse(w, r, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if errRes.Code == http.StatusNotFound {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
		return
	} else if errRes.Code == http.StatusBadRequest {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
		return
	} else if errRes.Code == http.StatusMethodNotAllowed {
		JsoneResponse(w, r, errRes.Msg, errRes.Code)
		return
	} else if errRes.Code == http.StatusInternalServerError {
		JsoneResponse(w, r, errRes.Msg, http.StatusInternalServerError)
		return
	} else if errRes.Code == http.StatusForbidden {
		JsoneResponse(w, r, errRes.Msg, http.StatusForbidden)
		return
	}
}
