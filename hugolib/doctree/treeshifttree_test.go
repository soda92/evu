package doctree_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/hugolib/doctree"
)

func TestTreeShiftTree(t *testing.T) {
	c := qt.New(t)

	tree := doctree.NewTreeShiftTree[string](0, 10)
	c.Assert(tree, qt.IsNotNil)
}
