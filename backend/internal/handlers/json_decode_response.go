package handlers

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(r *http.Request) *json.Decoder {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	return decode
}

func JsoneResponse(w http.ResponseWriter, message any, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(map[string]any{
		"message": message,
	})
	if err != nil {
		HandleError(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func NewEncoderJsone(w http.ResponseWriter) *json.Encoder {
	decode := json.NewEncoder(w)
	return decode
}