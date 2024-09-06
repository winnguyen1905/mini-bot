// Package messages contains all messages to be displayed.
package messages

import (
	"blitzbot/models"
	"blitzbot/pkg/i18n"
)

// Main returns main menu's message.
func Main(user *models.User) string {
	return i18n.Get(user.Language, "Welcome {firstname}!", map[string]interface{}{
		"firstname": user.FirstName,
	})
}
