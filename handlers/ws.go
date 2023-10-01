package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Establishing connection to ws")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error Upgrading to a WebSocket: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Con established")

	defer conn.Close()
	id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating new uuid")
		return
	}

	connections[id.String()] = conn
	log.Printf("Adding %s connection", id.String())

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading msg: %s", err)
			return
		}

		msg := string(p)
		log.Printf("Recibido %s", msg)
		if err := conn.WriteMessage(messageType, []byte("Respuesta: "+msg)); err != nil {
			log.Printf("Error writting msg: %s", err)
			return
		}
	}

}
