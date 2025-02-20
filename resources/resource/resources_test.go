package resource

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestResourcesMount(t *testing.T) {
	c := qt.New(t)
	c.Assert(true, qt.IsTrue)

	var m ResourceGetter
	var r Resources

	check := func(in, expect string) {
		c.Helper()
		r := m.Get(in)
		c.Assert(r, qt.Not(qt.IsNil))
		c.Assert(r.Name(), qt.Equals, expect)
	}

	checkNil := func(in string) {
		c.Helper()
		r := m.Get(in)
		c.Assert(r, qt.IsNil)
	}

	// Misc tests.
	r = Resources{
		testResource{name: "/foo/theme.css"},
	}

	m = r.Mount("/foo", ".")
	check("./theme.css", "/foo/theme.css")

	// Relative target.
	r = Resources{
		testResource{name: "/a/b/c/d.txt"},
		testResource{name: "/a/b/c/e/f.txt"},
		testResource{name: "/a/b/d.txt"},
		testResource{name: "/a/b/e.txt"},
	}

	m = r.Mount("/a/b/c", "z")
	check("z/d.txt", "/a/b/c/d.txt")
	check("z/e/f.txt", "/a/b/c/e/f.txt")

	m = r.Mount("/a/b", "")
	check("d.txt", "/a/b/d.txt")
	m = r.Mount("/a/b", ".")
	check("d.txt", "/a/b/d.txt")
	m = r.Mount("/a/b", "./")
	check("d.txt", "/a/b/d.txt")
	check("./d.txt", "/a/b/d.txt")

	m = r.Mount("/a/b", ".")
	check("./d.txt", "/a/b/d.txt")

	// Absolute target.
	m = r.Mount("/a/b/c", "/z")
	check("/z/d.txt", "/a/b/c/d.txt")
	check("/z/e/f.txt", "/a/b/c/e/f.txt")
	checkNil("/z/f.txt")

	m = r.Mount("/a/b", "/z")
	check("/z/c/d.txt", "/a/b/c/d.txt")
	check("/z/c/e/f.txt", "/a/b/c/e/f.txt")
	check("/z/d.txt", "/a/b/d.txt")
	checkNil("/z/f.txt")

	m = r.Mount("", "")
	check("/a/b/c/d.txt", "/a/b/c/d.txt")
	check("/a/b/c/e/f.txt", "/a/b/c/e/f.txt")
	check("/a/b/d.txt", "/a/b/d.txt")
	checkNil("/a/b/f.txt")

	m = r.Mount("/a/b", "/a/b")
	check("/a/b/c/d.txt", "/a/b/c/d.txt")
	check("/a/b/c/e/f.txt", "/a/b/c/e/f.txt")
	check("/a/b/d.txt", "/a/b/d.txt")
	checkNil("/a/b/f.txt")

	// Resources with relative paths.
	r = Resources{
		testResource{name: "a/b/c/d.txt"},
		testResource{name: "a/b/c/e/f.txt"},
		testResource{name: "a/b/d.txt"},
		testResource{name: "a/b/e.txt"},
		testResource{name: "n.txt"},
	}

	m = r.Mount("a/b", "z")
	check("z/d.txt", "a/b/d.txt")
	checkNil("/z/d.txt")
}

type testResource struct {
	Resource
	name string
}

func (r testResource) Name() string {
	return r.name
}

func (r testResource) NameNormalized() string {
	return r.name
}
