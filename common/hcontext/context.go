package hcontext

import "context"

// ContextDispatcher is a generic interface for setting and getting values from a context.
type ContextDispatcher[T any] interface {
	Set(ctx context.Context, value T) context.Context
	Get(ctx context.Context) T
}

// NewContextDispatcher creates a new ContextDispatcher with the given key.
func NewContextDispatcher[T any, R comparable](key R) ContextDispatcher[T] {
	return keyInContext[T, R]{
		id: key,
	}
}

type keyInContext[T any, R comparable] struct {
	zero T
	id   R
}

func (f keyInContext[T, R]) Get(ctx context.Context) T {
	v := ctx.Value(f.id)
	if v == nil {
		return f.zero
	}
	return v.(T)
}

func (f keyInContext[T, R]) Set(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, f.id, value)
}
