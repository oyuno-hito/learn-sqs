package controller

type Controllers struct {
	HealthController  HealthController
	MessageController MessageController
}

func NewControllers(
	healthController HealthController,
	messageController MessageController,
) Controllers {
	return Controllers{
		HealthController:  healthController,
		MessageController: messageController,
	}
}
