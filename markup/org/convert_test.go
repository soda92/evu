package org_test

import (
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config/testconfig"
	"github.com/spf13/afero"

	"github.com/gohugoio/hugo/markup/converter"
	"github.com/gohugoio/hugo/markup/org"

	qt "github.com/frankban/quicktest"
)

func TestConvert(t *testing.T) {
	c := qt.New(t)
	p, err := org.Provider.New(converter.ProviderConfig{
		Logger: loggers.NewDefault(),
		Conf:   testconfig.GetTestConfig(afero.NewMemMapFs(), nil),
	})
	c.Assert(err, qt.IsNil)
	conv, err := p.New(converter.DocumentContext{})
	c.Assert(err, qt.IsNil)
	b, err := conv.Convert(converter.RenderContext{Src: []byte("testContent")})
	c.Assert(err, qt.IsNil)
	c.Assert(string(b.Bytes()), qt.Equals, "<p>testContent</p>\n")
}
