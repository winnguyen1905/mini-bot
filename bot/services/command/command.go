// package command implements command request functionalities.
package command

import (
	service_interface "blitzbot/bot/services/interface"
	"blitzbot/database"
	"blitzbot/pkg/logger"
)

type cmd struct {
	db     database.IDatabase
	logger logger.Logger
}

// New returns a command service.
func New(db database.IDatabase, logger logger.Logger) service_interface.ICommand {
	return &cmd{
		db:     db,
		logger: logger,
	}
}
