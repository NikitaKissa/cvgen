package llm

import "context"

type Engine interface {
	LoadModel(ctx context.Context, path string) error
	Generate(ctx context.Context, req GenerateRequest) (string, error)
	Close() error
}
