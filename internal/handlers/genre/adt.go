package genre

import (
	"producer-portfolio-bot/internal/common"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const kulichMsg = `me: Prod/SoundDesign/Mix

https://www.instagram.com/reel/C6EWd4rIVjx/?igsh=aHV6ajBtZHAwY3pt`

const eklerMsg = `me: SoundDesign

https://www.instagram.com/reel/C4FmzN1IDf1/?igsh=MWl1MTBnOTh4cHplbA==`

const bezShorMsg = `me: Prod/SoundDesign/Mix

https://www.youtube.com/watch?v=2WlICm_XX9s&t=924s`

const brokkMsg = `me: Prod/Mix`

const trudMsg = `me: Prod

https://t.me/git_szfo/83`

const krispyMsg = `me: Prod/Edit/Mix`

func SendAdtAudio(bot *tgbotapi.BotAPI, chatID int64) {
	audioFiles := []struct {
		Path    string
		Caption string
	}{
		{"audio/Слой Куличи.mp3", kulichMsg},
		{"audio/Слой Эклеры.mp3", eklerMsg},
		{"audio/Интро к YouTube шоу Без Шор.mp3", bezShorMsg},
		{"audio/Ферма брокколей.mp3", brokkMsg},
		{"audio/Инспекция труда.mp3", trudMsg},
		{"audio/Криспи батончик.mp3", krispyMsg},
	}

	for i, audioFile := range audioFiles {
		common.SendAudioWithText(bot, chatID, audioFile.Path, audioFile.Caption, i == len(audioFiles)-1)
	}
}
