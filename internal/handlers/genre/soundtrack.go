package genre

import (
	"producer-portfolio-bot/internal/common"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

const begMsg = `Интро к Аудиопьессе Булгакова "Бег" 
me: Prod/SoundDesign/Edit/Mix

https://www.litres.ru/audiobook/mihail-bulgakov/beg-audiospektakl-68969814/`

func SendSoundtrackAudio(bot *tgbotapi.BotAPI, chatID int64) {
	audioFiles := []struct {
		Path    string
		Caption string
	}{
		{"audio/7 MINUTES.mp3", sevenMsg},
		{"audio/ANXIETY.mp3", anxietyMsg},
		{"audio/22_Глава_аудиоспектакля_Шарлотты_Хаберзак.mp3", habMsg},
		{"audio/BEG 0 INTRO.mp3", ""},
		{"audio/BEG 1 сон.mp3", ""},
		{"audio/BEG 2 сон.mp3", ""},
		{"audio/BEG 3 сон.mp3", ""},
		{"audio/BEG 4 сон.mp3", ""},
		{"audio/BEG 5 сон.mp3", ""},
		{"audio/BEG 6 сон.mp3", ""},
		{"audio/BEG 7 сон.mp3", begMsg},
	}

	for i, audioFile := range audioFiles {
		common.SendAudioWithText(bot, chatID, audioFile.Path, audioFile.Caption, i == len(audioFiles)-1)
	}
}
