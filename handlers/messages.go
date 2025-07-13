package handlers

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/convxz/ludo-bot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func HandleMessages(bot *tgbotapi.BotAPI, update tgbotapi.Update, db gorm.DB) {
	if update.Message.Text != "" {
		re := regexp.MustCompile(`^!(красное|черное)\s+(\d+)$`)
		if re.MatchString(update.Message.Text) {
			Roulette(bot, update, db)
		}
	}
}

func Roulette(bot *tgbotapi.BotAPI, update tgbotapi.Update, db gorm.DB) {
	re := regexp.MustCompile(`^!(красное|черное)\s+(\d+)$`)
	sub := re.FindStringSubmatch(update.Message.Text)
	color := sub[1]
	bet, _ := strconv.Atoi(sub[2])
	// добавить проверку на наличие денег на которые играть собираются

	balance := database.CheckBalance(int(update.Message.From.ID), db)
	if bet <= balance {
		animation := tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileID("CgACAgQAAxkBAANmaHL-4Hnr0ptFPNQtrMKcJXnK6FkAAvcEAAJl_L1QmxuJA4yUk182BA"))
		msgid, _ := bot.Send(animation)
		time.Sleep(2 * time.Second)
		del := tgbotapi.NewDeleteMessage(update.Message.From.ID, msgid.MessageID)
		bot.Request(del)

		options := []string{"красное", "черное"}
		randomChoice := options[rand.Intn(len(options))]
		if color == randomChoice {
			// функция для добавления денег на счет в зависимости от ставки
			msg := tgbotapi.NewMessage(update.Message.From.ID, "win. your balance: "+strconv.Itoa(balance+bet))
			database.ChangeBalance(int(update.Message.From.ID), balance+bet, db)
			bot.Send(msg)

		} else {
			// функция для отнимания денег со счета в зависимости от ставки
			msg := tgbotapi.NewMessage(update.Message.From.ID, "lose. your balance: "+strconv.Itoa(balance-bet))
			database.ChangeBalance(int(update.Message.From.ID), balance-bet, db)
			bot.Send(msg)
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.From.ID, "you have not a lot money")
		bot.Send(msg)
	}

}
