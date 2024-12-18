package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeJson(r *http.Request) *json.Decoder {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	defer r.Body.Close()
	return decode
}

func JsoneResponse(w http.ResponseWriter, r *http.Request, message any, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(map[string]any{
		"message": message,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func NewEncoderJsone(w http.ResponseWriter) *json.Encoder {
	decode := json.NewEncoder(w)
	return decode
}
