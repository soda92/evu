package hugio

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestHasBytesWriter(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c := qt.New((t))

	neww := func() (*HasBytesWriter, io.Writer) {
		var b bytes.Buffer

		h := &HasBytesWriter{
			Patterns: []*HasBytesPattern{
				{Pattern: []byte("__foo")},
			},
		}

		return h, io.MultiWriter(&b, h)
	}

	rndStr := func() string {
		return strings.Repeat("ab cfo", r.Intn(33))
	}

	for i := 0; i < 22; i++ {
		h, w := neww()
		fmt.Fprint(w, rndStr()+"abc __foobar"+rndStr())
		c.Assert(h.Patterns[0].Match, qt.Equals, true)

		h, w = neww()
		fmt.Fprint(w, rndStr()+"abc __f")
		fmt.Fprint(w, "oo bar"+rndStr())
		c.Assert(h.Patterns[0].Match, qt.Equals, true)

		h, w = neww()
		fmt.Fprint(w, rndStr()+"abc __moo bar")
		c.Assert(h.Patterns[0].Match, qt.Equals, false)
	}

	h, w := neww()
	fmt.Fprintf(w, "__foo")
	c.Assert(h.Patterns[0].Match, qt.Equals, true)
}
