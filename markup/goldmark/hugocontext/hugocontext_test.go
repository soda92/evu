package hugocontext

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestWrap(t *testing.T) {
	c := qt.New(t)

	b := []byte("test")

	c.Assert(Wrap(b, 42), qt.Equals, "{{__hugo_ctx pid=42}}\ntest\n{{__hugo_ctx/}}\n")
}

func BenchmarkWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wrap([]byte("test"), 42)
	}
}
