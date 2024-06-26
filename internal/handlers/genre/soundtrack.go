package genre

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const sevenMsg = `me: Prod/Rec/Edit/Mix/Master

https://youtu.be/B6e3_r97ckc?si=RE0d3XFTk-sNrMeF`

const anxietyMsg = `me: Prod/Rec/Edit/Mix/Master`

const habMsg = `me: Prod/SoundDesign/Voice&MusicEdit/Mix

1 часть:
https://www.litres.ru/audiobook/sharlotta-haberzak-18181113/ne-otkryvat-malipusechki-chast-1-70259830/

2 часть:
https://www.litres.ru/audiobook/sharlotta-haberzak-18181113/ne-otkryvat-malipusechki-chast-2-70259944/

3 часть:
https://www.litres.ru/audiobook/sharlotta-haberzak-18181113/ne-otkryvat-malipusechki-chast-3-70259998/`

func SendSoundtrackAudio(bot *tgbotapi.BotAPI, chatID int64) {
	audioFiles := []struct {
		Path    string
		Caption string
	}{
		{"audio/7 MINUTES.mp3", sevenMsg},
		{"ANXIETY.mp3", anxietyMsg},
		{"22_Глава_аудиоспектакля_Шарлотты_Хаберзак.mp3", habMsg},
	}

	for _, audioFile := range audioFiles {
		sendAudioWithText(bot, chatID, audioFile.Path, audioFile.Caption)
	}
}
