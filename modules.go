package main

import (
	"fmt"

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
