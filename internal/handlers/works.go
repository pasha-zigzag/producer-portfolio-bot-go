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
			tgbotapi.NewInlineKeyboardButtonData("Prod/Mixing", "MIXING"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Саундтреки", "SOUNDTRACK"),
			tgbotapi.NewInlineKeyboardButtonData("Adt", "AD"),
		),
	)

	// Создание и отправка сообщения с инлайн клавиатурой
	msg := tgbotapi.NewMessage(chatID, worksMessage)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)
}
