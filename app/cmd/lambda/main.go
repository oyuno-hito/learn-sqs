package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"learn-sqs/app/pkg/database"
	"learn-sqs/app/service/lambda/config/di"
	"log"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatalf(err.Error())
	}

	handler := di.Wire(db)

	lambda.Start(handler.Do)
}
