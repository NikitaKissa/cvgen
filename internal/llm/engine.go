package llm

import "context"

type GenerateRequest struct {
	Prompt      string
	MaxTokens   int
	Temperature float32
}

type Engine interface {
	LoadModel(ctx context.Context, path string) error
	Generate(ctx context.Context, req GenerateRequest) (<-chan string, error)
	Close() error
}
