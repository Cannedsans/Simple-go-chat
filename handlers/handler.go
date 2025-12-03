package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
    Username string `json:"username"`
    Message  string `json:"message"`
}



var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)


func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer conn.Close()
	clients[conn] = true

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			delete(clients, conn)
			break
		}

		broadcast <- msg
	}
}
