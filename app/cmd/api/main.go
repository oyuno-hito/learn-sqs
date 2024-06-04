package main

import (
	"github.com/gin-gonic/gin"
	"learn-sqs/app/pkg/database"
	"learn-sqs/app/service/api/config/di"
	"log"
)

func main() {
	router := gin.Default()

	db, err := database.Init()
	if err != nil {
		log.Fatalf(err.Error())
	}

	controllers := di.Wire(db)

	router.GET("/health", controllers.HealthController.GET)
	router.POST("/messages", controllers.MessageController.Post)

	_ = router.Run(":80")
}
