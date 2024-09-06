// package user implements user functionalities.
package user

import (
	service_interface "blitzbot/bot/services/interface"
	"blitzbot/database"
	"blitzbot/models"
	"blitzbot/pkg/logger"
	"context"
	"fmt"
)

type user struct {
	db     database.IDatabase
	logger logger.Logger
}

// New returns a new user service.
func New(db database.IDatabase, logger logger.Logger) service_interface.IUser {
	return &user{
		db:     db,
		logger: logger,
	}
}

// Create create a new user.
func (u *user) Create(ctx context.Context, user *models.User) error {
	if u.IsExist(ctx, user) {
		return nil
	}

	err := u.db.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Get returns a user's full information.
func (u *user) Get(ctx context.Context, user *models.User) (*models.User, error) {
	if !u.IsExist(ctx, user) {
		return nil, fmt.Errorf("user don't exist: %d", user.TelegramID)
	}

	user, err := u.db.GetUser(ctx, user.TelegramID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Language returns user's language.
func (u *user) Language(ctx context.Context, user *models.User) (string, error) {
	user, err := u.Get(ctx, user)
	if err != nil {
		return "", err
	}

	return user.Language, nil
}

// ChooseLanguage sets user's language.
func (u *user) ChooseLanguage(ctx context.Context, lang string, user *models.User) error {
	err := u.db.UpdateUserLanguage(ctx, lang, user.TelegramID)
	if err != nil {
		return err
	}

	return nil
}

// IsExist checks if user is already a member.
func (u *user) IsExist(ctx context.Context, user *models.User) bool {
	_, err := u.db.GetUser(ctx, user.TelegramID)

	return err == nil
}
