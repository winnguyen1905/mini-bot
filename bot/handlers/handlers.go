// package handlers handles all incoming requests.
package handlers

import (
	"blitzbot/bot/services"
	"blitzbot/database"
	"blitzbot/pkg/logger"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handler handles all services.
type Handler struct {
	db       database.IDatabase
	logger   logger.Logger
	services *services.Service
}

type IHandler interface {
	Callback(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.MessageConfig
	Command(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.EditMessageTextConfig
	Message(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.EditMessageTextConfig
}

func New(db database.IDatabase, logger logger.Logger) *Handler {
	return &Handler{
		db:       db,
		logger:   logger,
		services: services.New(db, logger),
	}
}
