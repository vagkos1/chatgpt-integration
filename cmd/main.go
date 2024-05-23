package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vagkos1/chatgpt-integration/internal/application"
	"github.com/vagkos1/chatgpt-integration/internal/infrastructure/chatgpt"
	"github.com/vagkos1/chatgpt-integration/internal/infrastructure/config"
	"github.com/vagkos1/chatgpt-integration/internal/infrastructure/logger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.LoadConfig()

	log := logger.NewLogger()

	chatgptClient := chatgpt.NewClient(cfg.ChatGPTAPIKey)

	chatgptService := application.NewChatGPTService(chatgptClient)

	systemContent := "You are a poetic assistant, skilled in explaining complex programming concepts with creative flair."
	userContent := "Compose a poem that explains the concept of recursion in programming."

	response, err := chatgptService.GetResponse(systemContent, userContent)
	if err != nil {
		log.Error("Error sending prompt", err)
		return
	}

	fmt.Println("ChatGPT Response:", response)
	log.Info("Successfully received response from ChatGPT")
}
