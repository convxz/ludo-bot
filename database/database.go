package database

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Balance int64
	ID      int64 `gorm:"primaryKey"`
}

func Init() gorm.DB {
	dsn := "host=localhost user=go password=go1234 dbname=go port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	return *db
}

// функция для добавления персонажа в бд.
// return 1 -> пользователь создан.
// return 2 -> пользователь уже существует.
func FirstTouch(update tgbotapi.Update, db gorm.DB) int {
	id := update.Message.From.ID
	var user User
	res := db.First(&user, id)
	if res.Error != nil {
		user = User{
			ID:      update.Message.From.ID,
			Balance: 2500,
		}
		db.Create(&user)
		return 1
	}
	return 2
}

// возвращает баланс по id
func CheckBalance(id int, db gorm.DB) int {
	var user User
	db.First(&user, id)
	return int(user.Balance)
}

// принимает id и balance, чтобы заменить баланс по id
func ChangeBalance(id int, balance int, db gorm.DB) {
	db.Model(&User{}).Where("id = ?", id).Update("balance", balance)
}
