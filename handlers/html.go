package handlers

import (
	"html/template"
	"net/http"
)

func renderHTML(w http.ResponseWriter, tmpl string, data interface{}) {
	htmlTmlp, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = htmlTmlp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
