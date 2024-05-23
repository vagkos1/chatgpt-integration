package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vagkos1/chatgpt-integration/internal/domain"
)

type Client struct {
	apiKey string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Ensure Client implements the ChatGPT interface
var _ domain.ChatGPT = (*Client)(nil)

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) SendPrompt(systemContent, userContent string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	requestBody, err := json.Marshal(ChatGPTRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role:    "system",
				Content: systemContent,
			},
			{
				Role:    "user",
				Content: userContent,
			},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var chatGPTResponse ChatGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		return "", err
	}

	if len(chatGPTResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from ChatGPT")
	}

	return chatGPTResponse.Choices[0].Message.Content, nil
}
