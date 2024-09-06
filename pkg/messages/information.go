// Package messages contains all messages to be displayed.
package messages

import (
	"blitzbot/models"
	"blitzbot/pkg/i18n"
)

// Information returns user's information message.
func Information(user *models.User) string {
	return i18n.Get(user.Language, "Telegram ID: {telegramID}\nUsername: {username}\nFirstname: {firstname}\nLastname: {lastname}\nLanguage: {language}", map[string]interface{}{
		"telegramID": user.TelegramID,
		"username":   user.Username,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"language":   user.Language,
	})
}
