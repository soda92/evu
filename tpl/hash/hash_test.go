// Package hash provides non-cryptographic hash functions for template use.
package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/spf13/cast"
)

func TestXxHash(t *testing.T) {
	t.Parallel()
	c := qt.New(t)

	ns := New()

	h, err := ns.XxHash("The quick brown fox jumps over the lazy dog")
	c.Assert(err, qt.IsNil)
	// Facit: https://asecuritysite.com/encryption/xxhash?val=The%20quick%20brown%20fox%20jumps%20over%20the%20lazy%20dog
	c.Assert(h, qt.Equals, "0b242d361fda71bc")
}

func BenchmarkXxHash(b *testing.B) {
	const inputSmall = "The quick brown fox jumps over the lazy dog"
	inputLarge := strings.Repeat(inputSmall, 100)

	runBench := func(name, input string, b *testing.B, fn func(v any)) {
		b.Run(fmt.Sprintf("%s_%d", name, len(input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(input)
			}
		})
	}

	ns := New()
	fnXxHash := func(v any) {
		_, err := ns.XxHash(v)
		if err != nil {
			panic(err)
		}
	}

	fnFNv32a := func(v any) {
		_, err := ns.FNV32a(v)
		if err != nil {
			panic(err)
		}
	}

	// Copied from the crypto tpl/crypto package,
	// just to have something to compare the above with.
	fnMD5 := func(v any) {
		conv, err := cast.ToStringE(v)
		if err != nil {
			panic(err)
		}

		hash := md5.Sum([]byte(conv))
		_ = hex.EncodeToString(hash[:])
	}

	for _, input := range []string{inputSmall, inputLarge} {
		runBench("xxHash", input, b, fnXxHash)
		runBench("mdb5", input, b, fnMD5)
		runBench("fnv32a", input, b, fnFNv32a)
	}
}
