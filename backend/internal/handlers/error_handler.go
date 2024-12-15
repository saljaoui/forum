package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errsResponse struct {
	Code int `json:"code"`
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	var errRes errsResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&errRes)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if errRes.Code == http.StatusNotFound {
		JsoneResponse(w, r, "404 Not Found: The requested resource could not be located", errRes.Code)
	} else if errRes.Code == http.StatusBadRequest {
		JsoneResponse(w, r, "400 Bad Request: The server could not understand your request.", errRes.Code)
	} else if errRes.Code == http.StatusMethodNotAllowed {
		JsoneResponse(w, r, "405 Method Not Allowed: The requested HTTP method is not supported for this resource.", errRes.Code)
	} else {
		JsoneResponse(w, r, "500 Internal Server Error: The server encountered an unexpected condition.", http.StatusInternalServerError)
	}
}
