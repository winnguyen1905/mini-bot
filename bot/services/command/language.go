// package command implements command request functionalities.
package command

import (
	"context"

	"blitzbot/models"
	"blitzbot/pkg/keyboards"
	"blitzbot/pkg/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Language shows a select language message.
func (m *cmd) Language(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Language(user)
	msg.ReplyMarkup = keyboards.Language()

	return msg
}
