// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"blitzbot/pkg/i18n"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BackToMain returns back to main keyboard.
func BackToMain(lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(i18n.Get(lang, "Back to main"), "backToMain"),
		),
	)
}
