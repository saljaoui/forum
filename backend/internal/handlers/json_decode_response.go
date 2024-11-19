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

func JsoneResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": message,
	})
}
