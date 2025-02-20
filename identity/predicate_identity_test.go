// Package provides ways to identify values in Hugo. Used for dependency tracking etc.
package identity

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestGlobIdentity(t *testing.T) {
	c := qt.New(t)

	gid := NewGlobIdentity("/a/b/*")

	c.Assert(isNotDependent(gid, StringIdentity("/a/b/c")), qt.IsFalse)
	c.Assert(isNotDependent(gid, StringIdentity("/a/c/d")), qt.IsTrue)
	c.Assert(isNotDependent(StringIdentity("/a/b/c"), gid), qt.IsTrue)
	c.Assert(isNotDependent(StringIdentity("/a/c/d"), gid), qt.IsTrue)
}

func isNotDependent(a, b Identity) bool {
	f := NewFinder(FinderConfig{})
	r := f.Contains(a, b, -1)
	return r == 0
}

func TestPredicateIdentity(t *testing.T) {
	c := qt.New(t)

	isDependent := func(id Identity) bool {
		return id.IdentifierBase() == "foo"
	}
	isDependency := func(id Identity) bool {
		return id.IdentifierBase() == "baz"
	}

	id := NewPredicateIdentity(isDependent, isDependency)

	c.Assert(id.IsProbablyDependent(StringIdentity("foo")), qt.IsTrue)
	c.Assert(id.IsProbablyDependent(StringIdentity("bar")), qt.IsFalse)
	c.Assert(id.IsProbablyDependent(id), qt.IsFalse)
	c.Assert(id.IsProbablyDependent(NewPredicateIdentity(isDependent, nil)), qt.IsFalse)
	c.Assert(id.IsProbablyDependency(StringIdentity("baz")), qt.IsTrue)
	c.Assert(id.IsProbablyDependency(StringIdentity("foo")), qt.IsFalse)
}
