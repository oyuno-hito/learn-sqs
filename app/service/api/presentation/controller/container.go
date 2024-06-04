package controller

type Controllers struct {
	HealthController HealthController
}

func NewControllers(healthController HealthController) Controllers {
	return Controllers{
		HealthController: healthController,
	}
}
