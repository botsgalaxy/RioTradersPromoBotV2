package main

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

var startText = `<i><b>Daily 3 - 6 GOLD Signals, Click Here 👇</b></i> `
var startInlineButton = gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{
			gotgbot.InlineKeyboardButton{Text: "Join Now ♐️", Url: "https://t.me/Riotraders"},
		},
	},
}
var startButton = gotgbot.ReplyKeyboardMarkup{
	Keyboard: [][]gotgbot.KeyboardButton{
		{
			gotgbot.KeyboardButton{Text: "Free Gold Signals"},
			gotgbot.KeyboardButton{Text: "Free Forex Signals"},
		},
		{

			gotgbot.KeyboardButton{Text: "Free Stock Signals"},
			gotgbot.KeyboardButton{Text: "Free Crypto Signals"},
		},
	},
	ResizeKeyboard: true,
}

func registerHandlers(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", start))
	d.AddHandler(handlers.NewCommand("support", support))
	d.AddHandler(handlers.NewCommand("freesignals", freesignals))
	d.AddHandler(handlers.NewCommand("broadcast", broadcast))
	d.AddHandler(handlers.NewMessage(message.Text, messageResponse))
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {

	user := ctx.EffectiveUser
	chatId := ctx.EffectiveChat.Id
	b.SendMessage(chatId, fmt.Sprintf("Hey, <b>%s</b>", user.FirstName), &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: startButton,
	})

	_, err := b.SendMessage(chatId, startText, &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: startInlineButton,
	})

	DB.Create(&PromoBotUser{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		IsPremium:    user.IsPremium,
		LanguageCode: user.LanguageCode,
	})
	return err

}

func messageResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	text := ctx.EffectiveMessage.Text
	chatId := ctx.EffectiveChat.Id
	switch text {
	case "Free Gold Signals":
		msg := "Join <b>Free GOLD Signals</b>, click here 👇"
		sendButtonWithText(b, chatId, msg, "https://t.me/+Jil6ArNaTd4yZjA1")

	case "Free Stock Signals":
		msg := "Join <b>Free Stock Signals</b>, click here 👇"
		sendButtonWithText(b, chatId, msg, "https://t.me/+xcTc7pbCcwY2ZWY1")

	case "Free Crypto Signals":
		msg := "Join <b>Free Crypto Signals</b>, click here 👇"
		sendButtonWithText(b, chatId, msg, "https://t.me/+trZ-UNf6bSM0N2Y1")

	case "Free Forex Signals":
		msg := "Join <b>Free Forex Signals</b>, click here 👇"
		sendButtonWithText(b, chatId, msg, "https://t.me/+-qRWp_QWJgBjNzA9")

	default:

	}
	return nil
}

func sendButtonWithText(b *gotgbot.Bot, chatId int64, text string, buttonUrl string) {
	b.SendMessage(
		chatId,
		text,
		&gotgbot.SendMessageOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
					{
						gotgbot.InlineKeyboardButton{Text: "Join Now ♐️", Url: buttonUrl},
					},
				},
			},
		},
	)
}

func support(b *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveUser.Id
	_, err := b.SendMessage(
		chatId,
		`<b>Please contact us if you have any queries </b>`,
		&gotgbot.SendMessageOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
					{
						gotgbot.InlineKeyboardButton{Text: "✍️ Talk to Support", Url: "https://t.me/RioTradersSupport"},
					},
				},
			},
		},
	)

	return err

}

func freesignals(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	chatId := ctx.EffectiveChat.Id
	_, err := b.SendMessage(chatId, fmt.Sprintf("Hey, <b>%s</b>", user.FirstName), &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: startButton,
	})
	return err

}

func broadcast(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	if validateAdmin(user.Username) {
		if ctx.EffectiveMessage.ReplyToMessage != nil {
			var botUsers []PromoBotUser
			result := DB.Find(&botUsers)
			if result.Error != nil {
				return result.Error
			}
			counter := "🟢 Broadcast Started\n\n✅ Successful: %d \n❌ Failed: %d"
			msg, _ := b.SendMessage(user.Id, fmt.Sprintf(counter, 0, 0), nil)
			success := 0
			failed := 0

			for _, botUser := range botUsers {
				_, err := ctx.EffectiveMessage.ReplyToMessage.Forward(
					b,
					botUser.UserId,
					&gotgbot.ForwardMessageOpts{},
				)
				if err != nil {
					failed++
				} else {
					success++
				}

				time.Sleep(time.Second * 3)
				msg.EditText(b, fmt.Sprintf(counter, success, failed), nil)
			}

		} else {
			ctx.EffectiveMessage.Reply(
				b,
				"You have to reply to a message to broadcast!!!",
				&gotgbot.SendMessageOpts{},
			)
		}

	} else {
		ctx.EffectiveMessage.Reply(
			b,
			"You are not authorized to broadcast!!!",
			&gotgbot.SendMessageOpts{},
		)
	}

	return nil
}

func validateAdmin(username string) bool {
	for _, adminUsername := range ADMIN_USERNAMES {
		if username == adminUsername {
			return true
		}
	}
	return false
}
