package identity

import "sync"

// NewQuestion creates a new question with the given identity.
func NewQuestion[T any](id Identity) *Question[T] {
	return &Question[T]{
		Identity: id,
	}
}

// Answer takes a func that knows the answer.
// Note that this is a one-time operation,
// fn will not be invoked again it the question is already answered.
// Use Result to check if the question is answered.
func (q *Question[T]) Answer(fn func() T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.answered {
		return
	}

	q.fasit = fn()
	q.answered = true
}

// Result returns the fasit of the question (if answered),
// and a bool indicating if the question has been answered.
func (q *Question[T]) Result() (any, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.fasit, q.answered
}

// A Question is defined by its Identity and can be answered once.
type Question[T any] struct {
	Identity
	fasit T

	mu       sync.RWMutex
	answered bool
}
