//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"learn-sqs/app/service/lambda/handler"
	"learn-sqs/app/service/lambda/usecase"
)

func Wire(db *gorm.DB) handler.Handler {
	wire.Build(
		usecase.NewMessageUsecase,
		handler.NewHandler,
	)
	return handler.Handler{}
}
