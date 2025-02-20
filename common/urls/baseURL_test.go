package urls

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBaseURL(t *testing.T) {
	c := qt.New(t)

	b, err := NewBaseURLFromString("http://example.com/")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "http://example.com/")

	b, err = NewBaseURLFromString("http://example.com")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "http://example.com/")
	c.Assert(b.WithPathNoTrailingSlash, qt.Equals, "http://example.com")
	c.Assert(b.BasePath, qt.Equals, "/")

	p, err := b.WithProtocol("webcal://")
	c.Assert(err, qt.IsNil)
	c.Assert(p.String(), qt.Equals, "webcal://example.com/")

	p, err = b.WithProtocol("webcal")
	c.Assert(err, qt.IsNil)
	c.Assert(p.String(), qt.Equals, "webcal://example.com/")

	_, err = b.WithProtocol("mailto:")
	c.Assert(err, qt.Not(qt.IsNil))

	b, err = NewBaseURLFromString("mailto:hugo@rules.com")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "mailto:hugo@rules.com")

	// These are pretty constructed
	p, err = b.WithProtocol("webcal")
	c.Assert(err, qt.IsNil)
	c.Assert(p.String(), qt.Equals, "webcal:hugo@rules.com")

	p, err = b.WithProtocol("webcal://")
	c.Assert(err, qt.IsNil)
	c.Assert(p.String(), qt.Equals, "webcal://hugo@rules.com")

	// Test with "non-URLs". Some people will try to use these as a way to get
	// relative URLs working etc.
	b, err = NewBaseURLFromString("/")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "/")

	b, err = NewBaseURLFromString("")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "/")

	// BaseURL with sub path
	b, err = NewBaseURLFromString("http://example.com/sub")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "http://example.com/sub/")
	c.Assert(b.WithPathNoTrailingSlash, qt.Equals, "http://example.com/sub")
	c.Assert(b.BasePath, qt.Equals, "/sub/")
	c.Assert(b.BasePathNoTrailingSlash, qt.Equals, "/sub")

	b, err = NewBaseURLFromString("http://example.com/sub/")
	c.Assert(err, qt.IsNil)
	c.Assert(b.String(), qt.Equals, "http://example.com/sub/")
	c.Assert(b.HostURL(), qt.Equals, "http://example.com")
}
