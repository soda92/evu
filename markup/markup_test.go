package markup_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/config/testconfig"
	"github.com/gohugoio/hugo/markup"
	"github.com/gohugoio/hugo/markup/converter"
)

func TestConverterRegistry(t *testing.T) {
	c := qt.New(t)
	conf := testconfig.GetTestConfig(nil, nil)
	r, err := markup.NewConverterProvider(converter.ProviderConfig{Conf: conf})

	c.Assert(err, qt.IsNil)
	c.Assert("goldmark", qt.Equals, r.GetMarkupConfig().DefaultMarkdownHandler)

	checkName := func(name string) {
		p := r.Get(name)
		c.Assert(p, qt.Not(qt.IsNil))
		c.Assert(p.Name(), qt.Equals, name)
	}

	c.Assert(r.Get("foo"), qt.IsNil)
	c.Assert(r.Get("markdown").Name(), qt.Equals, "goldmark")

	checkName("goldmark")
	checkName("asciidocext")
	checkName("rst")
	checkName("pandoc")
	checkName("org")
}
