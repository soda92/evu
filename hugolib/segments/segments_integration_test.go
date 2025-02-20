package segments_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/hugolib"
)

func TestSegments(t *testing.T) {
	files := `
-- hugo.toml --
baseURL = "https://example.org/"
renderSegments = ["docs"]
[languages]
[languages.en]
weight = 1
[languages.no]
weight = 2
[languages.nb]
weight = 3
[segments]
[segments.docs]
[[segments.docs.includes]]
kind = "{home,taxonomy,term}"
[[segments.docs.includes]]
path = "{/docs,/docs/**}"
[[segments.docs.excludes]]
path = "/blog/**"
[[segments.docs.excludes]]
lang = "n*"
output = "rss"
[[segments.docs.excludes]]
output = "json"
-- layouts/_default/single.html --
Single: {{ .Title }}|{{ .RelPermalink }}|
-- layouts/_default/list.html --
List: {{ .Title }}|{{ .RelPermalink }}|
-- content/docs/_index.md --
-- content/docs/section1/_index.md --
-- content/docs/section1/page1.md --
---
title: "Docs Page 1"
tags: ["tag1", "tag2"]
---
-- content/blog/_index.md --
-- content/blog/section1/page1.md --
---
title: "Blog Page 1"
tags: ["tag1", "tag2"]
---
`

	b := hugolib.Test(t, files)
	b.Assert(b.H.Configs.Base.RootConfig.RenderSegments, qt.DeepEquals, []string{"docs"})

	b.AssertFileContent("public/docs/section1/page1/index.html", "Docs Page 1")
	b.AssertFileExists("public/blog/section1/page1/index.html", false)
	b.AssertFileExists("public/index.html", true)
	b.AssertFileExists("public/index.xml", true)
	b.AssertFileExists("public/no/index.html", true)
	b.AssertFileExists("public/no/index.xml", false)
}
