package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const contactsMessage = `+7-996-316-07-09 (RU)
Telegram https://t.me/mbrody`

func showContacts(bot *tgbotapi.BotAPI, chatID int64) {
	// Создание клавиатуры
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("К портфолио"),
		),
	)
	keyboard.OneTimeKeyboard = true
	keyboard.ResizeKeyboard = true

	// Создание и отправка сообщения с клавиатурой
	msg := tgbotapi.NewMessage(chatID, contactsMessage)
	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}
