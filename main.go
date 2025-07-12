package main

import (
	// "database/sql"
	"fmt"
	"log"
	"os"

	"github.com/convxz/ludo-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
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
	// con := "host=localhost port=5432 user=go password=go1234 dbname=go sslmode=disable"
	// db, err := sql.Open("postgres", con)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("db inited")

	// create table db
	// cdb := `
	// CREATE TABLE IF NOT EXISTS userswid (
	// 	id BIGINT,
	// 	message varchar(100)
	// )
	// `
	// db.Exec(cdb)

	// обновления раз в 60 секунд
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// прост повторяет за мной пока
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				handlers.HandleCommands(bot, update)
			} else {
				handlers.HandleMessages(bot, update)
			}

		}
	}
}
