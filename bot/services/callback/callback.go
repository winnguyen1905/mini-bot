// package callback implements callback request functionalities.
package callback

import (
	service_interface "blitzbot/bot/services/interface"
	"blitzbot/database"
	"blitzbot/pkg/logger"
)

type call struct {
	db     database.IDatabase
	logger logger.Logger
}

// New returns a callback service.
func New(db database.IDatabase, logger logger.Logger) service_interface.ICallback {
	return &call{
		db:     db,
		logger: logger,
	}
}
