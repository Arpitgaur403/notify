package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients map[int]*websocket.Conn
}

var WS = Hub{
	Clients: make(map[int]*websocket.Conn),
}

func Register(userID int, conn *websocket.Conn) {
	WS.Clients[userID] = conn
	log.Println("User connected:", userID)
}

func Send(userID int, message string) {
	conn, exists := WS.Clients[userID]
	if !exists {
		log.Println("User offline:", userID)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Send error:", err)
	}
}
