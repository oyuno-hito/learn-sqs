//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"learn-sqs/app/service/api/presentation/controller"
)

func Wire(db *gorm.DB) controller.Controllers {
	wire.Build(
		controller.NewControllers,
		controller.NewHealthController,
	)
	return controller.Controllers{}
}
