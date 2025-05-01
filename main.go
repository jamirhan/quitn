package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var shotsNum = 3
var intervalDuration = 3

var bot *tgbotapi.BotAPI
var chatID = int64(5448179159)

func sleepFor(d time.Duration) {
	// time.Sleep(d)
	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("will sleep for %s", d.String())))
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Panic("BOT_TOKEN is not set")
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	now := time.Now()

	todayAt13 := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		13, 0, 0, 0,
		now.Location(),
	)

	if time.Now().Hour() > 13 {
		bot.Send(tgbotapi.NewMessage(chatID, "bot exited unsuccessfully, manual action required, current time: "+time.Now().Format(time.RFC3339)))
		return
	}

	for {
		sleepFor(time.Until(todayAt13))
		for range shotsNum {
			sleepHours := rand.Intn(intervalDuration)
			sleepMinutes := rand.Intn(60)
			sleepFor(time.Duration(sleepHours)*time.Hour + time.Duration(sleepMinutes)*time.Minute)
			bot.Send(tgbotapi.NewMessage(chatID, "shot"))
		}
		time.Sleep(10 * time.Second)
	}
}
