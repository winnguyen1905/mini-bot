// Package models contain entities that are essential for the app.
package models

import (
	"time"
)

// User represents user table structure in a database.
type User struct {
	ID         int64     `bson:"id"`
	TelegramID int64     `bson:"telegram_id,omitempty"`
	Username   string    `bson:"username"`
	FirstName  string    `bson:"firstname,omitempty"`
	LastName   string    `bson:"lastname"`
	JoinedAt   time.Time `bson:"joined_at,omitempty"`
	Language   string    `bson:"language"`
}
