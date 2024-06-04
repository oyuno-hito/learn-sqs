// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"learn-sqs/app/service/api/presentation/controller"
)

// Injectors from wire.go:

func Wire() controller.Controllers {
	healthController := controller.NewHealthController()
	controllers := controller.NewControllers(healthController)
	return controllers
}