//go:build extended

package resources_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/htesting/hqt"
	"github.com/gohugoio/hugo/media"
)

func TestImageResizeWebP(t *testing.T) {
	c := qt.New(t)

	_, image := fetchImage(c, "sunrise.webp")

	c.Assert(image.MediaType(), qt.Equals, media.Builtin.WEBPType)
	c.Assert(image.RelPermalink(), qt.Equals, "/a/sunrise.webp")
	c.Assert(image.ResourceType(), qt.Equals, "image")
	exif := image.Exif()
	c.Assert(exif, qt.Not(qt.IsNil))
	c.Assert(exif.Tags["Copyright"], qt.Equals, "Bjørn Erik Pedersen")
	c.Assert(exif.Lat, hqt.IsSameFloat64, 36.59744166666667)
	c.Assert(exif.Long, hqt.IsSameFloat64, -4.50846)
	c.Assert(exif.Date.IsZero(), qt.Equals, false)

	resized, err := image.Resize("123x")
	c.Assert(err, qt.IsNil)
	c.Assert(image.MediaType(), qt.Equals, media.Builtin.WEBPType)
	c.Assert(resized.RelPermalink(), qt.Equals, "/a/sunrise_hu_a1deb893888915d9.webp")
	c.Assert(resized.Width(), qt.Equals, 123)
}
