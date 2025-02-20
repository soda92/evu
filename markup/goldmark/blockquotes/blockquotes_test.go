package blockquotes

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestResolveBlockQuoteAlert(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	tests := []struct {
		input    string
		expected blockQuoteAlert
	}{
		{
			input:    "[!NOTE]",
			expected: blockQuoteAlert{typ: "note"},
		},
		{
			input:    "[!FaQ]",
			expected: blockQuoteAlert{typ: "faq"},
		},
		{
			input:    "[!NOTE]+",
			expected: blockQuoteAlert{typ: "note", sign: "+"},
		},
		{
			input:    "[!NOTE]-",
			expected: blockQuoteAlert{typ: "note", sign: "-"},
		},
		{
			input:    "[!NOTE] This is a note",
			expected: blockQuoteAlert{typ: "note", title: "This is a note"},
		},
		{
			input:    "[!NOTE]+ This is a note",
			expected: blockQuoteAlert{typ: "note", sign: "+", title: "This is a note"},
		},
		{
			input:    "[!NOTE]+ This is a title\nThis is not.",
			expected: blockQuoteAlert{typ: "note", sign: "+", title: "This is a title"},
		},
		{
			input:    "[!NOTE]\nThis is not.",
			expected: blockQuoteAlert{typ: "note"},
		},
	}

	for i, test := range tests {
		c.Assert(resolveBlockQuoteAlert("<p>"+test.input+"</p>"), qt.Equals, test.expected, qt.Commentf("Test %d", i))
	}
}
