package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	database "learn-sqs/app/pkg/database/model"
	"net/http"
)

type MessageController struct {
	db *gorm.DB
}

type PostMessageRequest struct {
	Message string `json:"message"`
}

func NewMessageController(db *gorm.DB) MessageController {
	return MessageController{db: db}
}

func (c MessageController) Post(ctx *gin.Context) {
	var req PostMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := c.saveMessage(req.Message)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c MessageController) saveMessage(message string) error {
	model := database.Message{Text: message}
	return c.db.Create(&model).Error
}
