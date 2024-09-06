package service_interface

import (
	"blitzbot/models"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback implements callback functionalities.
type ICallback interface {
	Menu(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig
	Help(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig
	Information(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig
	Language(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig
}

// Command implements command functionalities.
type ICommand interface {
	Language(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
}

// Message implements message functionalities.
type IMessage interface {
	Wrong(ctx context.Context, msg tgbotapi.EditMessageTextConfig, user *models.User) tgbotapi.EditMessageTextConfig
}

// Account implements user functionalities.
type IUser interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, user *models.User) (*models.User, error)
	Language(ctx context.Context, user *models.User) (string, error)
	ChooseLanguage(ctx context.Context, lang string, user *models.User) error
	IsExist(ctx context.Context, user *models.User) bool
}
