// Package bot handles bot activities.
package handlers

import (
	"context"
	"time"

	"blitzbot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command handles command requests.
func (h *Handler) Command(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.MessageConfig {
	// deleteLastMessage(tgbot, update.Message.Chat.ID, update.Message.MessageID)

	user := &models.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		LastName:   update.Message.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.Command() == "start" {
		if h.services.User.IsExist(ctx, user) {
			msg = h.services.Command.Language(ctx, msg, user)
		} else {
			user.JoinedAt = time.Now()

			err := h.services.User.Create(ctx, user)
			if err != nil {
				h.logger.Error(err)
			}

			msg = h.services.Command.Language(ctx, msg, user)
		}
	}

	return msg
}
