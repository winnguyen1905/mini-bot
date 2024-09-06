// Package messages contains all messages to be displayed.
package messages

import (
	"blitzbot/models"
	"blitzbot/pkg/i18n"
)

// language returns a choosing language message.
func Language(user *models.User) string {
	return i18n.Get(user.Language, "Choose a language:")
}
