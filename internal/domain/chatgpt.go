package domain

type ChatGPT interface {
	SendPrompt(prompt string) (string, error)
}
