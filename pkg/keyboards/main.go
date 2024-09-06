// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"blitzbot/pkg/i18n"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Main returns maun menu's keyboard.
func Main(lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				i18n.Get(lang, "Your Information"),
				"information",
			),
			tgbotapi.NewInlineKeyboardButtonData(
				i18n.Get(lang, "Help"),
				"help",
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				i18n.Get(lang, "Choosing another language"),
				"language",
			),
		),
	)
}
