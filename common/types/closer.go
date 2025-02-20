package types

import "sync"

type Closer interface {
	Close() error
}

// CloserFunc is a convenience type to create a Closer from a function.
type CloserFunc func() error

func (f CloserFunc) Close() error {
	return f()
}

type CloseAdder interface {
	Add(Closer)
}

type Closers struct {
	mu sync.Mutex
	cs []Closer
}

func (cs *Closers) Add(c Closer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.cs = append(cs.cs, c)
}

func (cs *Closers) Close() error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	for _, c := range cs.cs {
		c.Close()
	}

	cs.cs = cs.cs[:0]

	return nil
}
