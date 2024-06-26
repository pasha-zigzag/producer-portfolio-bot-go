package handlers

import (
	"log"
	"producer-portfolio-bot/internal/handlers/genre"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	var text string
	if msg.IsCommand() {
		text = msg.Command()
	} else {
		text = msg.Text
	}

	switch text {
	case "start":
		handleStart(bot, msg)
	case "К портфолио":
		showWorks(bot, msg.Chat.ID)
	default:
		handleUnknown(bot, msg.Chat.ID)
	}
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	switch callbackQuery.Data {
	case "MY_WORKS":
		showWorks(bot, callbackQuery.From.ID)
	case "CONTACT_ME":
		showContacts(bot, callbackQuery.From.ID)
	case "MIXING":
		genre.SendMixingAudio(bot, callbackQuery.From.ID)
	case "SOUNDTRACK":
		genre.SendSoundtrackAudio(bot, callbackQuery.From.ID)
	case "AD":
		genre.SendAdtAudio(bot, callbackQuery.From.ID)
	default:
		handleUnknown(bot, callbackQuery.Message.Chat.ID)
	}

	// Отправка ответа на callback запрос
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Не удалось отправить callback: %v", err)
	}
}

func handleUnknown(bot *tgbotapi.BotAPI, chatId int64) {
	response := tgbotapi.NewMessage(chatId, "Неизвестная команда.")
	bot.Send(response)
}
