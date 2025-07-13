package main

import (
	"fmt"
	"log"
	"os"

	"github.com/convxz/ludo-bot/database"
	"github.com/convxz/ludo-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot_token := os.Getenv("bot_token")

	// init bot
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("bot created")

	// init db
	db := database.Init()

	// обновления раз в 60 секунд
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// прост повторяет за мной пока
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				handlers.HandleCommands(bot, update, db)
			} else {
				handlers.HandleMessages(bot, update)
			}

		}
	}
}
