package handler

import (
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Portal System"

	// Parsing templates from files
	tmp := template.Must(template.ParseFiles("./templates/index.html"))
	tmp.Execute(w, data)
}
