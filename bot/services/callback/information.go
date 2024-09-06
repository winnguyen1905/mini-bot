// package callback implements callback request functionalities.
package callback

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Information shows user's information.
func (c *call) Information(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig {
	msg.Text = messages.Information(user)
	msg.ReplyMarkup.InlineKeyboard = keyboards.BackToMain(user.Language).InlineKeyboard

	return msg
}
