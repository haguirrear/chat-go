package handlers

import (
	"net/http"
)

type PageData struct {
	Title   string
	Heading string
	Message string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Chat app",
		Heading: "Star chatting Now",
		Message: "Start",
	}

	renderHTML(w, "templates/chat.html", data)
}
