package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aleksjjk/hookah-bot/internal/config"
	"github.com/aleksjjk/hookah-bot/internal/handlers"
	"github.com/aleksjjk/hookah-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
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
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL не установлен")
	}

	// Настройка и установка вебхука
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Panicf("Ошибка при создании вебхука: %v", err)
	}

	_, err = bot.Request(webhookConfig)
	if err != nil {
		log.Panicf("Ошибка при установке вебхука: %v", err)
	}

	// Проверка статуса вебхука
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Panicf("Ошибка при получении информации о вебхуке: %v", err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	//Инициализация обработчика
	handler := handlers.NewHandler(bot, cfg.AdminChatID)

	// Настройка обновлений
	updates := bot.ListenForWebhook("/")

	// Запуск HTTP-сервера в отдельной горутине
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()
	// Основной цикл обработки обновлений
	for update := range updates {
		handler.HandleUpdate(update)
	}
}
