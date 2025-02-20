package images_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestAutoOrient(t *testing.T) {
	files := `
-- hugo.toml --
-- assets/rotate270.jpg --
sourcefilename: ../testdata/exif/orientation6.jpg
-- layouts/index.html --
{{ $img := resources.Get "rotate270.jpg" }}
W/H original: {{ $img.Width }}/{{ $img.Height }}
{{ $rotated := $img.Filter images.AutoOrient }}
W/H rotated: {{ $rotated.Width }}/{{ $rotated.Height }}
`

	b := hugolib.Test(t, files)
	b.AssertFileContent("public/index.html", "W/H original: 80/40\n\nW/H rotated: 40/80")
}

// Issue 12733.
func TestOrientationEq(t *testing.T) {
	files := `
-- hugo.toml --
-- assets/rotate270.jpg --
sourcefilename: ../testdata/exif/orientation6.jpg
-- layouts/index.html --
{{ $img := resources.Get "rotate270.jpg" }}
{{ $orientation := $img.Exif.Tags.Orientation }}
Orientation: {{ $orientation }}|eq 6: {{ eq $orientation 6 }}|Type: {{ printf "%T" $orientation }}|
`

	b := hugolib.Test(t, files)
	b.AssertFileContent("public/index.html", "Orientation: 6|eq 6: true|")
}
