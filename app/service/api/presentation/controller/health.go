package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func NewHealthController() HealthController {
	return HealthController{}
}

func (c HealthController) GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "service is healthy."})
}
