package llama

import (
	"context"
	"fmt"
	"sync"

	"github.com/NikitaKissa/cvgen/internal/llm"
)

type Llama struct {
	mu sync.Mutex

	modelPath string

	model   llamaModel
	ctx     llamaContext
	sampler llamaSampler

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

func New() llm.Engine {
	return &Llama{}
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

	model, err := loadModel(path)
	if err != nil {
		return err
	}

	llamaCtx, err := createContext(model)
	if err != nil {
		freeModel(model)
		return err
	}

	sampler, err := createSampler()
	if err != nil {
		freeContext(llamaCtx)
		freeModel(model)
		return err
	}

	l.model = model
	l.ctx = llamaCtx
	l.sampler = sampler
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

	return generate(ctx, l.model, l.ctx, l.sampler, req)
}

func (l *Llama) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return nil
	}

	if l.sampler.ptr != 0 {
		freeSampler(l.sampler)
		l.sampler = llamaSampler{}
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
