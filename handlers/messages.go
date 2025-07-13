package handlers

import (
	"math/rand"
	"regexp"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Text != "" {
		re := regexp.MustCompile(`^!(красное|черное)\s+(\d+)$`)
		if re.MatchString(update.Message.Text) {
			game(bot, update)
		}
	}
}

func game(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	re := regexp.MustCompile(`^!(красное|черное)\s+(\d+)$`)
	sub := re.FindStringSubmatch(update.Message.Text)
	color := sub[1]
	// number := sub[2]
	// добавить проверку на наличие денег на которые играть собираются

	animation := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID("CgACAgQAAxkBAANmaHL-4Hnr0ptFPNQtrMKcJXnK6FkAAvcEAAJl_L1QmxuJA4yUk182BA"))
	msgid, _ := bot.Send(animation)
	time.Sleep(3 * time.Second)
	del := tgbotapi.NewDeleteMessage(update.Message.From.ID, msgid.MessageID)
	bot.Request(del)

	options := []string{"красное", "черное"}
	randomChoice := options[rand.Intn(len(options))]
	if color == randomChoice {
		// функция для добавления денег на счет в зависимости от ставки
	} else {
		// функция для отнимания денег со счета в зависимости от ставки
	}
}
