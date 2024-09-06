// Package bot handles bot activities.
package handlers

import (
	"context"

	"blitzbot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message handles message requests.
func (h *Handler) Message(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.EditMessageTextConfig {

	user := &models.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		LastName:   update.Message.From.LastName,
	}

	msg := tgbotapi.NewEditMessageTextAndMarkup(update.Message.Chat.ID, update.Message.MessageID, "", tgbotapi.InlineKeyboardMarkup{})

	var err error
	user.Language, err = h.services.User.Language(ctx, user)
	if err != nil {
		h.logger.Error(err)
	}

	msg = h.services.Message.Wrong(ctx, msg, user)

	return msg
}
