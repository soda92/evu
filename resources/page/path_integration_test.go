package page_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

// Issue 4926
// Issue 8232
// Issue 12342
func TestHashSignInPermalink(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ['section','rss','sitemap','taxonomy']
[permalinks]
s1 = '/:section/:slug'
-- layouts/_default/list.html --
{{ range site.Pages }}{{ .RelPermalink }}|{{ end }}
-- layouts/_default/single.html --
{{ .Title }}
-- content/s1/p1.md --
---
title: p#1
tags: test#tag#
---
-- content/s2/p#2.md --
---
title: p#2
---
`

	b := hugolib.Test(t, files)

	b.AssertFileExists("public/s1/p#1/index.html", true)
	b.AssertFileExists("public/s2/p#2/index.html", true)
	b.AssertFileExists("public/tags/test#tag#/index.html", true)

	b.AssertFileContentExact("public/index.html", "/|/s1/p%231/|/s2/p%232/|/tags/test%23tag%23/|")
}
