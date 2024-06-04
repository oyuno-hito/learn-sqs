//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"learn-sqs/app/service/api/presentation/controller"
)

func Wire() controller.Controllers {
	wire.Build(
		controller.NewControllers,
		controller.NewHealthController,
	)
	return controller.Controllers{}
}
