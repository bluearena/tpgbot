/*
Package main starts bot like server application, connets to Telegram API and writes
logs in local Postgres database
*/
package main

import (
	"log"
	"strings"
	"tpgbot/analyzer"
	"tpgbot/config"
	"tpgbot/db"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

//Main starts bot with params
func main() {
	botStarter(60)
}

//botStater run bot
func botStarter(timeout int) {
	lg := db.Logs{}
	bot, err := tgbotapi.NewBotAPI(config.BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		//Prepeare log content and add it to PG database
		lg.Login = update.Message.From.UserName
		lg.Chat = update.Message.Chat.Title
		lg.Text = update.Message.Text
		lg.Time = time.Now().String()
		//Run log writer in new threat
		go db.AddMessageToLog(&lg)
		//Text analyzer and reactions
		if strings.Compare(lg.Text, "") != 0 {
			answer := analyzer.GetAnswer(lg.Text)
			if strings.Compare(answer, "") != 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
