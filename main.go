package main

import (
	"log"

	"producer-portfolio-bot/config"
	"producer-portfolio-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatalf("Не удалось создать бота: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go handlers.HandleMessage(bot, update.Message)
		}
	}
}
