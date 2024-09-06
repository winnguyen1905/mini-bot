package services

import (
	"blitzbot/bot/services/callback"
	"blitzbot/bot/services/command"
	service_interface "blitzbot/bot/services/interface"
	"blitzbot/bot/services/message"
	"blitzbot/bot/services/user"
	"blitzbot/database"
	"blitzbot/pkg/logger"
)

type Service struct {
	logger   logger.Logger
	Callback service_interface.ICallback
	Command  service_interface.ICommand
	Message  service_interface.IMessage
	User     service_interface.IUser
}

func New(db database.IDatabase, logger logger.Logger) *Service {
	return &Service{
		logger:   logger,
		Callback: callback.New(db, logger),
		Command:  command.New(db, logger),
		Message:  message.New(db, logger),
		User:     user.New(db, logger),
	}
}
