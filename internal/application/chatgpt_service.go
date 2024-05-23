package application

import (
	"github.com/vagkos1/chatgpt-integration/internal/domain"
)

type ChatGPTService struct {
	chatgptClient domain.ChatGPT
}

func NewChatGPTService(chatgptClient domain.ChatGPT) *ChatGPTService {
	return &ChatGPTService{chatgptClient: chatgptClient}
}

func (s *ChatGPTService) GetResponse(prompt string) (string, error) {
	return s.chatgptClient.SendPrompt(prompt)
}
