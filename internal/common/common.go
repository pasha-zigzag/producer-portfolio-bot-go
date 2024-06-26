package common

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendAudioWithText(bot *tgbotapi.BotAPI, chatID int64, audioPath, caption string, withKeyboard bool) {
	audio := tgbotapi.NewAudio(chatID, tgbotapi.FilePath(audioPath))

	audio.Caption = caption

	if withKeyboard {
		var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Мои работы", "MY_WORKS"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Как связаться со мной", "CONTACT_ME"),
			),
		)
		audio.ReplyMarkup = inlineKeyboard
	}
	if _, err := bot.Send(audio); err != nil {
		log.Println("Failed to send audio:", err)
	}
}
