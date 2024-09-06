package database

import (
	"blitzbot/models"
	"context"
)

type IDatabase interface {
	IUserRepo
	Close(ctx context.Context) error
}

type IUserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, telegramID int64) (*models.User, error)
	UpdateUserLanguage(ctx context.Context, lang string, telegramID int64) error
}
