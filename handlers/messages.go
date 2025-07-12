package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Text == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "it's not a text")
		bot.Send(msg)
	}
}
