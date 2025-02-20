// Package provides ways to identify values in Hugo. Used for dependency tracking etc.
package identity_test

import (
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func BenchmarkFinder(b *testing.B) {
	m1 := identity.NewManager("")
	m2 := identity.NewManager("")
	m3 := identity.NewManager("")
	m1.AddIdentity(
		testIdentity{"base", "id1", "", "pe1"},
		testIdentity{"base2", "id2", "eq1", ""},
		m2,
		m3,
	)

	b4 := testIdentity{"base4", "id4", "", ""}
	b5 := testIdentity{"base5", "id5", "", ""}

	m2.AddIdentity(b4)

	f := identity.NewFinder(identity.FinderConfig{})

	b.Run("Find one", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := f.Contains(b4, m1, -1)
			if r == 0 {
				b.Fatal("not found")
			}
		}
	})

	b.Run("Find none", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := f.Contains(b5, m1, -1)
			if r > 0 {
				b.Fatal("found")
			}
		}
	})
}
