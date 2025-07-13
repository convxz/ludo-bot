package handlers

import (
	"github.com/convxz/ludo-bot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func HandleCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update, db gorm.DB) {
	switch update.Message.Command() {
	case "start":
		database.FirstTouch(update, db)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "привет! и пусть удача всегда будет с тобой :)")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "unknown comm")
		bot.Send(msg)
	}
}
