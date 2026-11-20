package llm

type GenerateRequest struct {
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
