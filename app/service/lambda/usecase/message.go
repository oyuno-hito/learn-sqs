package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"gorm.io/gorm"
	database "learn-sqs/app/pkg/database/model"
)

type EventUsecase struct {
	db *gorm.DB
}

func NewMessageUsecase(db *gorm.DB) *EventUsecase {
	return &EventUsecase{
		db: db,
	}
}

func (u *EventUsecase) Save(ctx context.Context, event events.SQSEvent) error {
	var message database.Message

	rawMessage := event.Records[0].Body
	fmt.Println(rawMessage)
	if err := json.Unmarshal([]byte(rawMessage), &message); err != nil {
		return err
	}

	model := database.Event{
		ID:   message.ID,
		Text: message.Text,
	}

	return u.db.Create(&model).Error
}
