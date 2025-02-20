package hugolib

import "testing"

// #10794
func TestFragmentsAndToCCrossSiteAccess(t *testing.T) {
	files := `
-- hugo.toml --
baseURL = "https://example.com"
disableKinds = ["taxonomy", "term", "home"]
defaultContentLanguage = "en"
defaultContentLanguageInSubdir = true
[languages]
[languages.en]
weight = 1
[languages.fr]
weight = 2
-- content/p1.en.md --
---
title: "P1"
outputs: ["HTML", "JSON"]
---

## Heading 1 EN

-- content/p1.fr.md --
---
title: "P1"
outputs: ["HTML", "JSON"]
---

## Heading 1 FR
-- layouts/_default/single.html --
HTML
-- layouts/_default/single.json --
{{ $secondSite := index .Sites 1 }}
{{ $p1 := $secondSite.GetPage "p1" }}
ToC: {{ $p1.TableOfContents }}
Fragments : {{ $p1.Fragments.Identifiers }}



	
`

	b := NewIntegrationTestBuilder(
		IntegrationTestConfig{
			TxtarString: files,
			T:           t,
		},
	).Build()

	b.AssertFileContent("public/en/p1/index.html", "HTML")
	b.AssertFileContent("public/en/p1/index.json", "ToC: <nav id=\"TableOfContents\">\n  <ul>\n    <li><a href=\"#heading-1-fr\">Heading 1 FR</a></li>\n  </ul>\n</nav>\nFragments : [heading-1-fr]")
}

// Issue #10866
func TestTableOfContentsWithIncludedMarkdownFile(t *testing.T) {
	files := `
-- hugo.toml --
baseURL = "https://example.com"
disableKinds = ["taxonomy", "term", "home"]
-- content/p1.md --
---
title: "P1"
---

## Heading P1 1
{{% include "p2" %}}

-- content/p2.md --
---
title: "P2"
---

### Heading P2 1
### Heading P2 2

-- layouts/shortcodes/include.html --
{{ with site.GetPage (.Get 0) }}{{ .RawContent }}{{ end }}
-- layouts/_default/single.html --
Fragments: {{ .Fragments.Identifiers }}|


	
`

	b := NewIntegrationTestBuilder(
		IntegrationTestConfig{
			TxtarString: files,
			T:           t,
		},
	).Build()

	b.AssertFileContent("public/p1/index.html", "Fragments: [heading-p1-1 heading-p2-1 heading-p2-2]|")
	b.AssertFileContent("public/p2/index.html", "Fragments: [heading-p2-1 heading-p2-2]|")
}
