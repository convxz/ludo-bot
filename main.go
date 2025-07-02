package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot_token := os.Getenv("bot_token")
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil {
		log.Panic(err)
	}

	// просто проверка
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// обновления раз в 60 секунд
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// прост повторяет за мной пока
	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "you said: "+update.Message.Text)
			bot.Send(msg)
		}
	}
}
