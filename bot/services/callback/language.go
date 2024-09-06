// package callback implements callback request functionalities.
package callback

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Language shows a select language menu.
func (m *call) Language(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig {
	msg.Text = messages.Language(user)
	msg.ReplyMarkup.InlineKeyboard = keyboards.Language().InlineKeyboard

	return msg
}
