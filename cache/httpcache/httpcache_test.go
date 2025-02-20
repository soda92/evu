package httpcache

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestGlobMatcher(t *testing.T) {
	c := qt.New(t)

	g := GlobMatcher{
		Includes: []string{"**/*.jpg", "**.png", "**/bar/**"},
		Excludes: []string{"**/foo.jpg", "**.css"},
	}

	p, err := g.CompilePredicate()
	c.Assert(err, qt.IsNil)

	c.Assert(p("foo.jpg"), qt.IsFalse)
	c.Assert(p("foo.png"), qt.IsTrue)
	c.Assert(p("foo/bar.jpg"), qt.IsTrue)
	c.Assert(p("foo/bar.png"), qt.IsTrue)
	c.Assert(p("foo/bar/foo.jpg"), qt.IsFalse)
	c.Assert(p("foo/bar/foo.css"), qt.IsFalse)
	c.Assert(p("foo.css"), qt.IsFalse)
	c.Assert(p("foo/bar/foo.css"), qt.IsFalse)
	c.Assert(p("foo/bar/foo.xml"), qt.IsTrue)
}
