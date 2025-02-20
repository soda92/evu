package sass

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestIsUnquotedCSSValue(t *testing.T) {
	c := qt.New(t)

	for _, test := range []struct {
		in  any
		out bool
	}{
		{"24px", true},
		{"1.5rem", true},
		{"10%", true},
		{"hsl(0, 0%, 100%)", true},
		{"calc(24px + 36px)", true},
		{"24xxx", true}, // a false positive.
		{123, true},
		{123.12, true},
		{"#fff", true},
		{"#ffffff", true},
		{"#ffffffff", false},
	} {
		c.Assert(isTypedCSSValue(test.in), qt.Equals, test.out)
	}
}
