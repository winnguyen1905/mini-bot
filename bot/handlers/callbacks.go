// Package bot handles bot activities.
package handlers

import (
	"context"
	"strings"

	"blitzbot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback handles callback requests.
func (h *Handler) Callback(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) tgbotapi.EditMessageTextConfig {
	user := &models.User{
		TelegramID: update.CallbackQuery.From.ID,
		Username:   update.CallbackQuery.From.UserName,
		FirstName:  update.CallbackQuery.From.FirstName,
		LastName:   update.CallbackQuery.From.LastName,
	}

	msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "", tgbotapi.InlineKeyboardMarkup{})

	var err error
	user.Language, err = h.services.User.Language(ctx, user)
	if err != nil {
		h.logger.Error(err)
	}

	if strings.HasPrefix(update.CallbackQuery.Data, "language:") {
		lang := strings.Split(update.CallbackQuery.Data, ":")[1]

		user.Language = lang
		err = h.services.User.ChooseLanguage(ctx, user.Language, user)
		msg = h.services.Callback.Menu(ctx, msg, user)
		if err != nil {
			h.logger.Error(err)
		}
	}

	switch update.CallbackQuery.Data {
	case "information":
		msg = h.services.Callback.Information(ctx, msg, user)
	case "help":
		msg = h.services.Callback.Help(ctx, msg, user)
	case "language":
		msg = h.services.Callback.Language(ctx, msg, user)
	case "back-to-main":
		msg = h.services.Callback.Menu(ctx, msg, user)
	default:
		msg = h.services.Callback.Menu(ctx, msg, user)
	}

	return msg
}
