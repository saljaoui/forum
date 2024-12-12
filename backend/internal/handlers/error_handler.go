package handlers

import (
	"net/http"
	"text/template"
)

func HandleError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("../../frontend/templates/err.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, struct {
		Msg  string
		Code int
	}{
		Msg:  msg,
		Code: code,
	})
}
