package page_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestTaxonomiesGetAndCount(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ['rss','sitemap']
[taxonomies]
author = 'authors'
-- layouts/_default/home.html --
John Smith count: {{ site.Taxonomies.authors.Count "John Smith" }}
Robert Jones count: {{ (site.Taxonomies.authors.Get "Robert Jones").Pages.Len }}
-- layouts/_default/single.html --
{{ .Title }}|
-- layouts/_default/list.html --
{{ .Title }}|
-- content/p1.md --
---
title: p1
authors: [John Smith,Robert Jones]
---
-- content/p2.md --
---
title: p2
authors: [John Smith]
---
`

	b := hugolib.Test(t, files)

	b.AssertFileContent("public/index.html",
		"John Smith count: 2",
		"Robert Jones count: 1",
	)
}

func TestTaxonomiesPage(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ['rss','section','sitemap']
[taxonomies]
tag = 'tags'
category = 'categories'
-- content/p1.md --
---
title: p1
tags: [tag-a]
---
-- layouts/_default/list.html --
{{- with site.Taxonomies.tags.Page }}{{ .RelPermalink }}{{ end }}|
{{- with site.Taxonomies.categories.Page }}{{ .RelPermalink }}{{ end }}|
-- layouts/_default/single.html --
{{ .Title }}
`

	b := hugolib.Test(t, files)

	b.AssertFileContent("public/index.html", "/tags/||")
}
