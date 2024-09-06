// package user implements message request functionalities.
package message

import (
	service_interface "blitzbot/bot/services/interface"
	"blitzbot/database"
	"blitzbot/pkg/logger"
)

type msg struct {
	db     database.IDatabase
	logger logger.Logger
}

// New returns a message service.
func New(db database.IDatabase, logger logger.Logger) service_interface.IMessage {
	return &msg{
		db:     db,
		logger: logger,
	}
}
