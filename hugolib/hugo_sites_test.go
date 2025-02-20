package hugolib

import "testing"

func TestSitesAndLanguageOrder(t *testing.T) {
	files := `
-- hugo.toml --
defaultContentLanguage = "fr"
defaultContentLanguageInSubdir = true
[languages]
[languages.en]
weight = 1
[languages.fr]
weight = 2
[languages.de]
weight = 3
-- layouts/index.html --
{{ $bundle := site.GetPage "bundle" }}
Bundle all translations: {{ range $bundle.AllTranslations }}{{ .Lang }}|{{ end }}$
Bundle translations: {{ range $bundle.Translations }}{{ .Lang }}|{{ end }}$
Site languages: {{ range site.Languages }}{{ .Lang }}|{{ end }}$
Sites: {{ range site.Sites }}{{ .Language.Lang }}|{{ end }}$
-- content/bundle/index.fr.md --
---
title: "Bundle Fr"
---
-- content/bundle/index.en.md --
---
title: "Bundle En"
---
-- content/bundle/index.de.md --
---
title: "Bundle De"
---
	
	`
	b := Test(t, files)

	b.AssertFileContent("public/en/index.html",
		"Bundle all translations: en|fr|de|$",
		"Bundle translations: fr|de|$",
		"Site languages: en|fr|de|$",
		"Sites: fr|en|de|$",
	)
}
