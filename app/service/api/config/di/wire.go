//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	sqsclient "learn-sqs/app/pkg/sqs"
	"learn-sqs/app/service/api/presentation/controller"
)

func Wire(db *gorm.DB, sqs *sqsclient.Sqs) controller.Controllers {
	wire.Build(
		controller.NewControllers,
		controller.NewHealthController,
		controller.NewMessageController,
	)
	return controller.Controllers{}
}
