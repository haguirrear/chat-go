package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// You can configure other options here if needed.
}

var connections = make(map[string]*websocket.Conn)

func broadcastMessage(message string) {
	log.Printf("Sending message to %d connections", len(connections))
	for _, conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("WebSocket write error: %s", err.Error())
		}
	}

	// a, ok := connections["asd"]
}
