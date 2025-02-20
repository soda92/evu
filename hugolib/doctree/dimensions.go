package doctree

const (
	// Language is currently the only dimension in the Hugo build matrix.
	DimensionLanguage DimensionFlag = 1 << iota
)

// Dimension is a row in the Hugo build matrix which currently has one value: language.
type Dimension [1]int

// DimensionFlag is a flag in the Hugo build matrix.
type DimensionFlag byte

// Has returns whether the given flag is set.
func (d DimensionFlag) Has(o DimensionFlag) bool {
	return d&o == o
}

// Set sets the given flag.
func (d DimensionFlag) Set(o DimensionFlag) DimensionFlag {
	return d | o
}

// Index returns this flag's index in the Dimensions array.
func (d DimensionFlag) Index() int {
	if d == 0 {
		panic("dimension flag not set")
	}
	return int(d - 1)
}
