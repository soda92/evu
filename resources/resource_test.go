package resources

import (
	"os"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestAtomicStaler(t *testing.T) {
	c := qt.New(t)

	type test struct {
		AtomicStaler
	}

	var v test

	c.Assert(v.StaleVersion(), qt.Equals, uint32(0))
	v.MarkStale()
	c.Assert(v.StaleVersion(), qt.Equals, uint32(1))
	v.MarkStale()
	c.Assert(v.StaleVersion(), qt.Equals, uint32(2))
}

func BenchmarkHashImage(b *testing.B) {
	f, err := os.Open("testdata/sunset.jpg")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := hashImage(f)
		if err != nil {
			b.Fatal(err)
		}
		f.Seek(0, 0)
	}
}
