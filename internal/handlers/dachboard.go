package handlers

import (
	"net/http"
	"text/template"
)

func Dachboard(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/home.html")
	tmp.Execute(w, nil)
}
