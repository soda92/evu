// Package diagrams provides template functions for generating diagrams.
package diagrams

import (
	"context"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/internal"
)

const name = "diagrams"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := &Namespace{
			d: d,
		}

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(cctx context.Context, args ...any) (any, error) { return ctx, nil },
		}

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
