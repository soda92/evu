// Package page provides template functions for accessing the current Page object,
// the entry level context for the current template.
package page

import (
	"context"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/tpl"

	"github.com/gohugoio/hugo/tpl/internal"
)

const name = "page"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ns := &internal.TemplateFuncsNamespace{
			Name: name,
			Context: func(ctx context.Context, args ...interface{}) (interface{}, error) {
				v := tpl.Context.Page.Get(ctx)
				if v == nil {
					// The multilingual sitemap does not have a page as its context.
					return nil, nil
				}

				return v.(page.Page), nil
			},
		}

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
