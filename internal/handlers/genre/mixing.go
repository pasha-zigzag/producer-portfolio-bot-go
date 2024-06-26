package genre

import (
	"producer-portfolio-bot/internal/common"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const myKingMsg = `Явление Music – Царь Мой (SINGLE 2023)
me: Prod/Rec/Edit
#IndiePop #worship #acoustic

https://www.youtube.com/watch?v=KK6TxwTugZI`

const myGodMsg = `Явление Music – Боже Мой (SINGLE 2023)
me: Prod/Rec/Edit/Mix
#IndiePop #worship

https://www.youtube.com/watch?v=lGc6H-8pfPc`

const youWinMsg = `Явление Music – Ты Победил (SINGLE 2024)
me: Prod/Rec/Edit/Mix/Master
#IndieRock #alternative 

https://www.youtube.com/watch?v=KQa9sTx4xus`

func SendMixingAudio(bot *tgbotapi.BotAPI, chatID int64) {
	audioFiles := []struct {
		Path    string
		Caption string
	}{
		{"audio/Явление Music - Царь мой.mp3", myKingMsg},
		{"audio/Явление Music - Боже Мой.mp3", myGodMsg},
		{"audio/Явление Music - Ты Победил.mp3", youWinMsg},
	}

	for i, audioFile := range audioFiles {
		common.SendAudioWithText(bot, chatID, audioFile.Path, audioFile.Caption, i == len(audioFiles)-1)
	}
}
