// Package site provides template functions for accessing the Site object.
package site

import (
	"context"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/resources/page"

	"github.com/gohugoio/hugo/tpl/internal"
)

const name = "site"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		s := page.WrapSite(d.Site)
		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(cctx context.Context, args ...any) (any, error) { return s, nil },
		}

		// We just add the Site as the namespace here. No method mappings.

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
