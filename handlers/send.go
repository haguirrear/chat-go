package handlers

import (
	"log"
	"net/http"
)

func SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Fail to parse form data", http.StatusInternalServerError)
		return
	}
	message := r.FormValue("message")

	log.Printf("Broadcasting msg: %q", message)
	broadcastMessage("Resp: " + message)

	renderHTML(w, "templates/message.html", message)
}
