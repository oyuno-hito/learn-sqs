package main

import (
	"github.com/gin-gonic/gin"
	"learn-sqs/app/service/api/config/di"
)

func main() {
	router := gin.Default()

	controllers := di.Wire()

	router.GET("/health", controllers.HealthController.GET)

	_ = router.Run(":80")
}
