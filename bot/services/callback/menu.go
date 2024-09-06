// package callback implements callback request functionalities.
package callback

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Menu shows main menu.
func (c *call) Menu(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig {
	msg.Text = messages.Main(user)
	msg.ReplyMarkup.InlineKeyboard = keyboards.Main(user.Language).InlineKeyboard

	return msg
}
