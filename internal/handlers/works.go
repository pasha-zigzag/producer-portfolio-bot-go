package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const worksMessage = "Выбери музыкальное направление"

// showWorks отображает меню с направлениями музыки
func showWorks(bot *tgbotapi.BotAPI, chatID int64) {
	// Создание инлайн клавиатуры
	var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Rock/Blues-Rock/Alternative", "ROCK"),
			tgbotapi.NewInlineKeyboardButtonData("Indie/Pop", "POP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Live/концерты", "LIVE"),
		),
	)

	// Создание и отправка сообщения с инлайн клавиатурой
	msg := tgbotapi.NewMessage(chatID, worksMessage)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)
}
