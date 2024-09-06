// Package messages contains all messages to be displayed.
package messages

import (
	"blitzbot/models"
	"blitzbot/pkg/i18n"
)

// Help returns a help message.
func Help(user *models.User) string {
	return i18n.Get(user.Language, "This is a clean architecture telegram bot in Golang!")
}
