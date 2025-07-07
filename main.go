package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	con := "host=localhost port=5432 user=go password=go1234 dbname=go sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("db inited")

	// create table db
	cdb := `
	CREATE TABLE IF NOT EXISTS userswid (
		id BIGINT,
		message varchar(100)
	)
	`
	db.Exec(cdb)

	// обновления раз в 60 секунд
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// прост повторяет за мной пока
	for update := range updates {
		if update.Message != nil {
			err := sendMessageToDB(*update.Message, db)
			id := update.Message.Chat.ID
			// msg := tgbotapi.NewMessage(id, "you said: "+text)

			if err != nil {
				log.Panic(err)
			}

			text, err := returnLastMessageUser(id, db)
			if err != nil {
				log.Panic(err)
			}

			msg := tgbotapi.NewMessage(id, text)
			bot.Send(msg)
		}
	}
}

func sendMessageToDB(message tgbotapi.Message, db *sql.DB) error {
	var ans int64
	id, text := message.Chat.ID, message.Text

	query := `INSERT INTO userswid (id, message) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, id, text).Scan(&ans)
	return err
}

func returnLastMessageUser(id int64, db *sql.DB) (string, error) {
	query := `SELECT message FROM userswid WHERE id = $1 ORDER BY RANDOM()`
	var msg string
	err := db.QueryRow(query, id).Scan(&msg)
	if err != nil {
		return "", err
	}
	return msg, nil

}
