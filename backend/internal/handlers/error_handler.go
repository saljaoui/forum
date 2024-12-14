package handlers

import (
	"html/template"
	"net/http"
)

// func HandleError(w http.ResponseWriter, r *http.Request, msg string, code int) {
// 	w.WriteHeader(code)
// 	tmpl, err := template.ParseFiles("../../frontend/templates/err.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, struct {
// 		Msg  string
// 		Code int
// 	}{
// 		Msg:  msg,
// 		Code: code,
// 	})
// 	fmt.Println("i'm here")
// 	//http.RedirectHandler("https://freshman.tech", http.StatusSeeOther)
// 	// http.Redirect(w, r, "../../frontend/templates/err.html", code)
// }

func HandleError(w http.ResponseWriter, r *http.Request, mes string, codes int) {
	 w.WriteHeader(codes)
	tmpl, err := template.ParseFiles("../../frontend/templates/err.html")
	if err != nil {
		http.Error(w, "Error loading error page", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, struct {
		Msg  string
		Code int
	}{
		Msg:  mes,
		Code: codes,
	})
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	return
}
