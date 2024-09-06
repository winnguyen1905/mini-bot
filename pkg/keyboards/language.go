// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"blitzbot/pkg/i18n"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	emoji "github.com/jayco/go-emoji-flag"
)

var buttons []tgbotapi.InlineKeyboardButton

func init() {
	locales := i18n.GetCurrentLocales()
	buttons = make([]tgbotapi.InlineKeyboardButton, len(locales))

	for i, locale := range locales {
		buttons[i] =
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%v %v", emoji.GetFlag(i18n.GetCountryCode(locale)), i18n.Get(locale, "choose_language")),
				"language:"+locale, // callback key, example: language:en-US
			)
	}
}

// Language returns choosing language keyboard.
func Language() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			buttons...,
		),
	)
}
