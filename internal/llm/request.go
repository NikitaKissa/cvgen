package llm

type Message struct {
	Role    string
	Content string
}

type GenerateRequest struct {
	Messages    []Message
	Prompt      string
	MaxTokens   int
	Temperature float32
}

const (
	DefaultMaxTokens   = 256
	DefaultTemperature = 0.7
)

func DefaultRequest(prompt string) GenerateRequest {
	return GenerateRequest{
		Prompt:      prompt,
		MaxTokens:   DefaultMaxTokens,
		Temperature: DefaultTemperature,
	}
}
