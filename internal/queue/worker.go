package queue

import (
	"log"

	"code.sli.ke/personal/notification/internal/db"
	"code.sli.ke/personal/notification/internal/websocket"
	"github.com/kataras/golog"
)

type NotificationJob struct {
	UserID  int
	Message string
}

var NotificationQueue = make(chan NotificationJob, 100)

func StartWorker() {
	go func() {
		for job := range NotificationQueue {
			_, err := db.DB.Exec(
				`INSERT INTO notifications(user_id, message) VALUES($1, $2)`,
				job.UserID,
				job.Message,
			)
			if err != nil {
				log.Println("DB Insert Error:", err)
				continue
			}
			websocket.Send(job.UserID, job.Message)
			golog.Warn("process queue job : ", job.UserID, " : ", job.Message)
			log.Println("Notification saved:", job.Message)

		}
	}()
}
