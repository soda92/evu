package identity

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestQuestion(t *testing.T) {
	c := qt.New(t)

	q := NewQuestion[int](StringIdentity("2+2?"))

	v, ok := q.Result()
	c.Assert(ok, qt.Equals, false)
	c.Assert(v, qt.Equals, 0)

	q.Answer(func() int {
		return 4
	})

	v, ok = q.Result()
	c.Assert(ok, qt.Equals, true)
	c.Assert(v, qt.Equals, 4)
}
