package pageparser

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func Test_lexIntroSection(t *testing.T) {
	t.Parallel()
	c := qt.New(t)
	for i, tt := range []struct {
		input                string
		expectItemType       ItemType
		expectSummaryDivider []byte
	}{
		{"{\"title\": \"JSON\"}\n", TypeFrontMatterJSON, summaryDivider},
		{"#+TITLE: ORG\n", TypeFrontMatterORG, summaryDividerOrg},
		{"+++\ntitle = \"TOML\"\n+++\n", TypeFrontMatterTOML, summaryDivider},
		{"---\ntitle: YAML\n---\n", TypeFrontMatterYAML, summaryDivider},
		// Issue 13152
		{"# ATX Header Level 1\n", tText, summaryDivider},
	} {
		errMsg := qt.Commentf("[%d] %v", i, tt.input)

		l := newPageLexer([]byte(tt.input), lexIntroSection, Config{})
		l.run()

		c.Assert(l.items[0].Type, qt.Equals, tt.expectItemType, errMsg)
		c.Assert(l.summaryDivider, qt.DeepEquals, tt.expectSummaryDivider, errMsg)

	}
}
