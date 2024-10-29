package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/aleksjjk/hookah-bot/internal/config"
	"github.com/aleksjjk/hookah-bot/internal/handlers"
	"github.com/aleksjjk/hookah-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация бота
	bot := utils.NewBot(cfg.BotToken)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Инициализация обработчика
	handler := handlers.NewHandler(bot, cfg.AdminChatID)

	// Настройка обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Основной цикл обработки обновлений
	for update := range updates {
		handler.HandleUpdate(update)
	}
}
