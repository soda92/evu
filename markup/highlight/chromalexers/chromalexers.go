package chromalexers

import (
	"sync"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
)

type lexersMap struct {
	lexers map[string]chroma.Lexer
	mu     sync.RWMutex
}

var lexerCache = &lexersMap{lexers: make(map[string]chroma.Lexer)}

// Get returns a lexer for the given language name, nil if not found.
// This is just a wrapper around chromalexers.Get that caches the result.
// Reasoning for this is that chromalexers.Get is slow in the case where the lexer is not found,
// which is a common case in Hugo.
func Get(name string) chroma.Lexer {
	lexerCache.mu.RLock()
	lexer, found := lexerCache.lexers[name]
	lexerCache.mu.RUnlock()

	if found {
		return lexer
	}

	lexer = lexers.Get(name)

	lexerCache.mu.Lock()
	lexerCache.lexers[name] = lexer
	lexerCache.mu.Unlock()

	return lexer
}
