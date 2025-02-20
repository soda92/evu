package hstring

import (
	"html/template"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/spf13/cast"
)

func TestRenderedString(t *testing.T) {
	c := qt.New(t)

	// Validate that it will behave like a string in Hugo settings.
	c.Assert(cast.ToString(HTML("Hugo")), qt.Equals, "Hugo")
	c.Assert(template.HTML(HTML("Hugo")), qt.Equals, template.HTML("Hugo"))
}
