package hstring

import (
	"html/template"

	"github.com/gohugoio/hugo/common/types"
)

var _ types.PrintableValueProvider = HTML("")

// HTML is a string that represents rendered HTML.
// When printed in templates it will be rendered as template.HTML and considered safe so no need to pipe it into `safeHTML`.
// This type was introduced as a wasy to prevent a common case of inifinite recursion in the template rendering
// when the `linkify` option is enabled with a common (wrong) construct like `{{ .Text | .Page.RenderString }}` in a hook template.
type HTML string

func (s HTML) String() string {
	return string(s)
}

func (s HTML) PrintableValue() any {
	return template.HTML(s)
}
