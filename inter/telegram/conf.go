package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramConfig struct {
	Bot    *tgbotapi.BotAPI
	Token  string
	ChatID int64
}

func NewTelegramConfig(token string, chatID int64) *TelegramConfig {
	return &TelegramConfig{
		Token:  token,
		ChatID: chatID,
	}
}
