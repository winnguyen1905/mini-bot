// Package bot handles bot activities.
package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"blitzbot/bot/handlers"
	"blitzbot/configs"
	"blitzbot/database"
	"blitzbot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type bot struct {
	handler *handlers.Handler
	logger  logger.Logger
}

type IBot interface {
	Start(ctx context.Context, cfg configs.App, debugMode bool) error
}

// New creates a new bot.
func New(db database.IDatabase, logger logger.Logger) IBot {
	return &bot{
		handler: handlers.New(db, logger),
		logger:  logger,
	}
}

// Start starts the bot.
func (b *bot) Start(ctx context.Context, cfg configs.App, debugMode bool) error {
	var err error
	b.logger.Info("Bot Starting...")

	tgbot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return err
	}

	go b.setupWebhookServer(ctx, tgbot)
	err = setWebhook(tgbot, cfg.WebhookURL)
	tgbot.Debug = debugMode

	b.logger.Info(fmt.Sprintf("Authorized on account %s", tgbot.Self.UserName))

	return err
}

func (b *bot) setupWebhookServer(ctx context.Context, tgbot *tgbotapi.BotAPI) {
	server := &http.Server{Addr: "0.0.0.0:3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		update, err := tgbot.HandleUpdate(r)
		if err != nil {
			errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(errMsg)
			return
		}

		// Only command is sending new message
		if update.Message != nil && update.Message.IsCommand() {
			_, err = tgbot.Send(b.handler.Command(ctx, tgbot, *update))
			if err != nil {
				b.logger.Error(err)
			}
			return
		}

		// Edit message for: callback, message
		var msg tgbotapi.EditMessageTextConfig

		if update.CallbackQuery != nil {
			msg = b.handler.Callback(ctx, tgbot, *update)
		} else if update.Message != nil {
			msg = b.handler.Message(ctx, tgbot, *update)
		} else {
			return
		}

		tgbotapi.WriteToHTTPResponse(w, msg)
	})

	// Run the server in a goroutine so it doesn't block the context listening.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log unexpected server error
			fmt.Printf("ListenAndServe(): %s\n", err)
		}
	}()

	// Listen for context cancellation
	<-ctx.Done()

	// Shutdown the server when context is canceled
	b.logger.Info("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		// Log shutdown error
		b.logger.Error("Server Shutdown Failed:%+v", err)
	} else {
		b.logger.Info("Server exited properly")
	}
}

func setWebhook(tgbot *tgbotapi.BotAPI, url string) error {
	wh, _ := tgbotapi.NewWebhook(url)

	_, err := tgbot.Request(wh)
	if err != nil {
		return err
	}

	info, err := tgbot.GetWebhookInfo()
	if info.LastErrorDate != 0 {
		err = fmt.Errorf("telegram callback failed: %s", info.LastErrorMessage)
	}

	return err
}
