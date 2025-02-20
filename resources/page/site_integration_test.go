package page_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

// Issue 12513
func TestPageSiteSitesDefault(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ['page','rss','section','sitemap','taxonomy','term']
defaultContentLanguage = 'de'
defaultContentLanguageInSubdir = true
[languages.en]
languageName = 'English'
weight = 1
[languages.de]
languageName = 'Deutsch'
weight = 2
-- layouts/index.html --
{{ .Site.Sites.Default.Language.LanguageName }}
`

	b := hugolib.Test(t, files)

	b.AssertFileContent("public/de/index.html", "Deutsch")
}
