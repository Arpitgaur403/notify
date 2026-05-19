package handler

import (
	"net/http"
	"strconv"

	ws "code.sli.ke/personal/notification/internal/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")

	userID, _ := strconv.Atoi(userIDStr)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	ws.Register(userID, conn)
}
