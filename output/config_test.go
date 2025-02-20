package output

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/media"
)

func TestDecodeConfig(t *testing.T) {
	c := qt.New(t)

	mediaTypes := media.Types{media.Builtin.JSONType, media.Builtin.XMLType}

	tests := []struct {
		name        string
		m           map[string]any
		shouldError bool
		assert      func(t *testing.T, name string, f Formats)
	}{
		{
			"Redefine JSON",
			map[string]any{
				"json": map[string]any{
					"baseName":    "myindex",
					"isPlainText": "false",
				},
			},
			false,
			func(t *testing.T, name string, f Formats) {
				msg := qt.Commentf(name)
				c.Assert(len(f), qt.Equals, len(DefaultFormats), msg)
				json, _ := f.GetByName("JSON")
				c.Assert(json.BaseName, qt.Equals, "myindex")
				c.Assert(json.MediaType, qt.Equals, media.Builtin.JSONType)
				c.Assert(json.IsPlainText, qt.Equals, false)
			},
		},
		{
			"Add XML format with string as mediatype",
			map[string]any{
				"MYXMLFORMAT": map[string]any{
					"baseName":  "myxml",
					"mediaType": "application/xml",
				},
			},
			false,
			func(t *testing.T, name string, f Formats) {
				c.Assert(len(f), qt.Equals, len(DefaultFormats)+1)
				xml, found := f.GetByName("MYXMLFORMAT")
				c.Assert(found, qt.Equals, true)
				c.Assert(xml.BaseName, qt.Equals, "myxml")
				c.Assert(xml.MediaType, qt.Equals, media.Builtin.XMLType)

				// Verify that we haven't changed the DefaultFormats slice.
				json, _ := f.GetByName("JSON")
				c.Assert(json.BaseName, qt.Equals, "index")
			},
		},
		{
			"Add format unknown mediatype",
			map[string]any{
				"MYINVALID": map[string]any{
					"baseName":  "mymy",
					"mediaType": "application/hugo",
				},
			},
			true,
			func(t *testing.T, name string, f Formats) {
			},
		},
	}

	for _, test := range tests {
		result, err := DecodeConfig(mediaTypes, test.m)
		msg := qt.Commentf(test.name)

		if test.shouldError {
			c.Assert(err, qt.Not(qt.IsNil), msg)
		} else {
			c.Assert(err, qt.IsNil, msg)
			test.assert(t, test.name, result.Config)
		}
	}
}
