package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"
)

var ADMIN_USERNAMES []string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	admins := os.Getenv("ADMIN_USERNAMES")
	ADMIN_USERNAMES = strings.Split(admins, ",")
	go migrateDatabase(dsn)
}

func main() {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)
	log.Printf("<< Registering Handlers >>")
	registerHandlers(dispatcher)

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)
	setBotMenu(b)
	updater.Idle()
}

func setBotMenu(b *gotgbot.Bot) {
	commands := []gotgbot.BotCommand{
		{Command: "start", Description: "start the bot"},
		{Command: "support", Description: "get support"},
		{Command: "freesignals", Description: "get free signals"},
	}
	ok, _ := b.SetMyCommands(commands, nil)
	if ok {
		log.Println("Succesfully updated bot commands!!!")
	}
}
