package llama

import (
	"context"
	"fmt"
	"sync"

	"github.com/NikitaKissa/cvgen/internal/llm"
)

type Llama struct {
	mu sync.Mutex

	modelPath    string
	chatTemplate string
	rawPrompt    bool

	model llamaModel
	ctx   llamaContext

	closed bool
}

type llamaModel struct {
	ptr uintptr
}

type llamaContext struct {
	ptr uintptr
}

type llamaSampler struct {
	ptr uintptr
}

type Option func(*Llama)

func WithChatTemplate(template string) Option {
	return func(l *Llama) {
		l.chatTemplate = template
	}
}

func WithRawPrompt() Option {
	return func(l *Llama) {
		l.rawPrompt = true
	}
}

func New(options ...Option) llm.Engine {
	l := &Llama{}
	for _, option := range options {
		option(l)
	}
	return l
}

func (l *Llama) LoadModel(ctx context.Context, path string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return fmt.Errorf("llama: engine is closed")
	}

	if l.model.ptr != 0 {
		return fmt.Errorf("llama: model already loaded")
	}
	if l.rawPrompt && l.chatTemplate != "" {
		return fmt.Errorf("llama: raw prompt mode cannot be combined with a chat template")
	}

	model, err := loadModel(path)
	if err != nil {
		return err
	}

	llamaCtx, err := createContext(model)
	if err != nil {
		freeModel(model)
		return err
	}

	l.model = model
	l.ctx = llamaCtx
	l.modelPath = path

	return nil
}

func (l *Llama) Generate(ctx context.Context, req llm.GenerateRequest) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return "", fmt.Errorf("llama: engine is closed")
	}

	if l.model.ptr == 0 {
		return "", fmt.Errorf("llama: model is not loaded")
	}

	return generate(ctx, l.model, l.ctx, l.chatTemplate, l.rawPrompt, req)
}

func (l *Llama) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return nil
	}

	if l.ctx.ptr != 0 {
		freeContext(l.ctx)
		l.ctx = llamaContext{}
	}

	if l.model.ptr != 0 {
		freeModel(l.model)
		l.model = llamaModel{}
	}

	l.closed = true

	return nil
}
