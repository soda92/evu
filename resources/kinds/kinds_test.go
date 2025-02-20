package kinds

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestKind(t *testing.T) {
	t.Parallel()
	c := qt.New(t)
	// Add tests for these constants to make sure they don't change
	c.Assert(KindPage, qt.Equals, "page")
	c.Assert(KindHome, qt.Equals, "home")
	c.Assert(KindSection, qt.Equals, "section")
	c.Assert(KindTaxonomy, qt.Equals, "taxonomy")
	c.Assert(KindTerm, qt.Equals, "term")

	c.Assert(GetKindMain("TAXONOMYTERM"), qt.Equals, KindTaxonomy)
	c.Assert(GetKindMain("Taxonomy"), qt.Equals, KindTaxonomy)
	c.Assert(GetKindMain("Page"), qt.Equals, KindPage)
	c.Assert(GetKindMain("Home"), qt.Equals, KindHome)
	c.Assert(GetKindMain("SEction"), qt.Equals, KindSection)

	c.Assert(GetKindAny("Page"), qt.Equals, KindPage)
	c.Assert(GetKindAny("Robotstxt"), qt.Equals, KindRobotsTXT)
}
