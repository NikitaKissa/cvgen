package llm

import "context"

type engineKey struct{}

func WithContext(ctx context.Context, engine Engine) context.Context {
	return context.WithValue(ctx, engineKey{}, engine)
}

func FromContext(ctx context.Context) (Engine, bool) {
	if engine, ok := ctx.Value(engineKey{}).(Engine); ok {
		return engine, true
	}
	return nil, false
}
