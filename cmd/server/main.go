package main

import (
	"code.sli.ke/personal/notification/internal/db"
	"code.sli.ke/personal/notification/internal/queue"
	"code.sli.ke/personal/notification/internal/route"
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

func main() {
	golog.Info("Main function started: ")
	db.Connect()
	r := gin.Default()
	route.Setup(r)
	queue.StartWorker()
	r.Run(":8080")
}
