package doctree

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestDimensionFlag(t *testing.T) {
	c := qt.New(t)

	var zero DimensionFlag
	var d DimensionFlag
	var o DimensionFlag = 1
	var p DimensionFlag = 12

	c.Assert(d.Has(o), qt.Equals, false)
	d = d.Set(o)
	c.Assert(d.Has(o), qt.Equals, true)
	c.Assert(d.Has(d), qt.Equals, true)
	c.Assert(func() { zero.Index() }, qt.PanicMatches, "dimension flag not set")
	c.Assert(DimensionLanguage.Index(), qt.Equals, 0)
	c.Assert(p.Index(), qt.Equals, 11)
}
