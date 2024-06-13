package handlers

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	responseText := `Привет! Я Михей, студийный звукорежиссер, аранжировщик и клавишник. Если тебе нужно:

~~Сделать саундтрек/саунд-дизайн к творческому, коммерческому видео или рекламе;

~~Хорошо записать, свести и сделать мастеринг песни (в том числе и концерта); Сделать качественную аудио-редакцию и монтаж (чистка/реставрация/тональная и ритмическая коррекция голосов и инструментов);

~~Написать аранжировку к песне 

~~Консультация по работе в Ableton Live (настройка проектов для студийной записи, лайвов, темплейты для разных задач); 

Ниже представлены некоторые из моих работ по написанию, записи, сведению и мастерингу музыки.`

	// Отправка текстового сообщения
	textMsg := tgbotapi.NewMessage(msg.Chat.ID, responseText)
	bot.Send(textMsg)

	// Путь к изображению
	photoPath := "photo/mbrody.jpg" // Убедитесь, что путь правильный

	// Отправка фото
	photoFile, err := os.Open(photoPath)
	if err != nil {
		log.Printf("Не удалось открыть файл: %v", err)
		return
	}
	defer photoFile.Close()

	// Создание инлайн клавиатуры
	var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мои работы", "MY_WORKS"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Как связаться со мной", "CONTACT_ME"),
		),
	)

	photo := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileReader{Name: "mbrody.jpg", Reader: photoFile})
	photo.ReplyMarkup = inlineKeyboard
	bot.Send(photo)
}
