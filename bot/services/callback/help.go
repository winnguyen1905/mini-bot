// package callback implements callback request functionalities.
package callback

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Help shows help message.
func (c *call) Help(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig {
	msg.Text = messages.Help(user)
	msg.ReplyMarkup.InlineKeyboard = keyboards.BackToMain(user.Language).InlineKeyboard

	return msg
}
