package resources_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestTransformCached(t *testing.T) {
	files := `
-- hugo.toml --
disableKinds = ["taxonomy", "term"]	
-- assets/css/main.css --
body {
	  background: #fff;
}
-- content/p1.md --
---
title: "P1"
---
P1.
-- content/p2.md --
---
title: "P2"
---
P2.
-- layouts/_default/list.html --
List.
-- layouts/_default/single.html --
{{ $css := resources.Get "css/main.css" | resources.Minify  }}
CSS: {{ $css.Content }}
`

	b := hugolib.Test(t, files)

	b.AssertFileContent("public/p1/index.html", "CSS: body{background:#fff}")
}
