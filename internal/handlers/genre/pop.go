package genre

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const indiePopMessage = "asd"

func SendIndiePopAudio(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, indiePopMessage)
	bot.Send(msg)

	// Отправка аудиофайлов
	// audioFiles := []string{"path/to/indie_pop_1.mp3", "path/to/indie_pop_2.mp3"} // Убедитесь, что пути правильные

	// for _, audioPath := range audioFiles {
	// 	// audio := tgbotapi.NewAudioUpload(chatID, audioPath)
	// 	// bot.Send(audio)
	// }
}
