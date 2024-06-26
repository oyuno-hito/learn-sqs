package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"learn-sqs/app/service/lambda/usecase"
)

type Handler struct {
	usecase *usecase.EventUsecase
}

func NewHandler(
	usecase *usecase.EventUsecase,
) Handler {
	return Handler{
		usecase: usecase,
	}
}

func (h Handler) Do(ctx context.Context, event events.SQSEvent) error {
	return h.usecase.Save(ctx, event)
}
