package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	database "learn-sqs/app/pkg/database/model"
	sqsclient "learn-sqs/app/pkg/sqs"
	"net/http"
)

type MessageController struct {
	db  *gorm.DB
	sqs *sqsclient.Sqs
}

type PostMessageRequest struct {
	Message string `json:"message"`
}

func NewMessageController(
	db *gorm.DB,
	sqs *sqsclient.Sqs,
) MessageController {
	return MessageController{
		db:  db,
		sqs: sqs,
	}
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

	err = c.sendMessage(ctx, req.Message)
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

func (c MessageController) sendMessage(ctx context.Context, message string) error {
	return c.sqs.SendMessage(ctx, message)
}
