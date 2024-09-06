// package user implements message request functionalities.
package message

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Wrong shows a wrong message.
func (m *msg) Wrong(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig {
	msg.Text = messages.Wrong(user)
	msg.ReplyMarkup.InlineKeyboard = keyboards.BackToMain(user.Language).InlineKeyboard

	return msg
}
