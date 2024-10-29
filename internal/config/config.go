package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	BotToken    string
	AdminChatID int64
}

func LoadConfig() *Config {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("BOT_TOKEN is not set")
	}

	adminChatIDStr := os.Getenv("ADMIN_CHAT_ID")
	if adminChatIDStr == "" {
		log.Panic("ADMIN_CHAT_ID is not set")
	}

	adminChatID, err := strconv.ParseInt(adminChatIDStr, 10, 64)
	if err != nil {
		log.Panic("Invalid ADMIN_CHAT_ID")
	}

	return &Config{
		BotToken:    botToken,
		AdminChatID: adminChatID,
	}
}
