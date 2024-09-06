// Package messages contains all messages to be displayed.
package messages

import (
	"blitzbot/models"
	"blitzbot/pkg/i18n"
)

// wrong returns wrong command or message.
func Wrong(user *models.User) string {
	return i18n.Get(user.Language, "Enter a valid message or command.")
}
