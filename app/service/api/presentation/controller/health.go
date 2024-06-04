package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	database "learn-sqs/app/pkg/database/model"
	"net/http"
)

type HealthController struct {
	// NOTE: アーキテクチャとしてイマイチだが単純な実装なので許容する
	db *gorm.DB
}

func NewHealthController(db *gorm.DB) HealthController {
	return HealthController{db: db}
}

func (c HealthController) GET(ctx *gin.Context) {
	health, err := c.dbHealthCheck()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "service is healthy, but database value is empty."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "service is unhealthy."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": health.Message})
}

func (c HealthController) dbHealthCheck() (database.Health, error) {
	health := new(database.Health)

	err := c.db.First(health).Error

	return *health, err
}
