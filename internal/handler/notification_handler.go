package handler

import (
	"net/http"

	"code.sli.ke/personal/notification/internal/db"
	"code.sli.ke/personal/notification/internal/models"
	"code.sli.ke/personal/notification/internal/queue"
	"github.com/gin-gonic/gin"
)

type NotificationReq struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

func CreateNotification(c *gin.Context) {
	var req NotificationReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "notification received",
		"data": req,
	})
	queue.NotificationQueue <- queue.NotificationJob{
		UserID:  req.UserID,
		Message: req.Message,
	}

	return
}

func GetNotifications(c *gin.Context) {
	rows, _ := db.DB.Query(
		"SELECT id, user_id, message, status, created_at, read_at FROM notifications",
	)

	var notifications []models.Notification

	for rows.Next() {
		var n models.Notification
		rows.Scan(&n.ID, &n.UserId, &n.Message, &n.Status, &n.CreatedAt, &n.ReadAt)
		notifications = append(notifications, n)
	}

	c.JSON(200, notifications)
}

func MarkAsRead(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(
		"UPDATE notifications SET status='READ', read_at=NOW() WHERE id=$1",
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "notification marked as read",
	})
}
