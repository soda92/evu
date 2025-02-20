package pagesfromdata

import "testing"

func BenchmarkHash(b *testing.B) {
	m := map[string]any{
		"foo":         "bar",
		"bar":         "foo",
		"stringSlice": []any{"a", "b", "c"},
		"intSlice":    []any{1, 2, 3},
		"largeText":   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit.",
	}

	bs := BuildState{}

	for i := 0; i < b.N; i++ {
		bs.hash(m)
	}
}
