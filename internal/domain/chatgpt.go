package domain

type ChatGPT interface {
	SendPrompt(systemContent, userContent string) (string, error)
}
