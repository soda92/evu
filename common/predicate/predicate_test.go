package predicate_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/common/predicate"
)

func TestAdd(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1

	c.Assert(p(1), qt.IsTrue)
	c.Assert(p(2), qt.IsFalse)

	neg := p.Negate()
	c.Assert(neg(1), qt.IsFalse)
	c.Assert(neg(2), qt.IsTrue)

	and := p.And(intP2)
	c.Assert(and(1), qt.IsFalse)
	c.Assert(and(2), qt.IsFalse)
	c.Assert(and(10), qt.IsTrue)

	or := p.Or(intP2)
	c.Assert(or(1), qt.IsTrue)
	c.Assert(or(2), qt.IsTrue)
	c.Assert(or(10), qt.IsTrue)
	c.Assert(or(11), qt.IsFalse)
}

func TestFilter(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1
	p = p.Or(intP2)

	ints := []int{1, 2, 3, 4, 1, 6, 7, 8, 2}

	c.Assert(p.Filter(ints), qt.DeepEquals, []int{1, 2, 1, 2})
	c.Assert(ints, qt.DeepEquals, []int{1, 2, 1, 2, 1, 6, 7, 8, 2})
}

func TestFilterCopy(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1
	p = p.Or(intP2)

	ints := []int{1, 2, 3, 4, 1, 6, 7, 8, 2}

	c.Assert(p.FilterCopy(ints), qt.DeepEquals, []int{1, 2, 1, 2})
	c.Assert(ints, qt.DeepEquals, []int{1, 2, 3, 4, 1, 6, 7, 8, 2})
}

var intP1 = func(i int) bool {
	if i == 10 {
		return true
	}
	return i == 1
}

var intP2 = func(i int) bool {
	if i == 10 {
		return true
	}
	return i == 2
}
