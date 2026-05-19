package route

import (
	"code.sli.ke/personal/notification/internal/handler"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/notify", handler.CreateNotification)
	r.GET("/notify", handler.GetNotifications)
	r.GET("/notify/:id/read", handler.MarkAsRead)
	r.GET("/ws", handler.WebSocketHandler)
}
