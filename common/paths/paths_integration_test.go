package paths_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestRemovePathAccents(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ["taxonomy", "term"]
defaultContentLanguage = "en"
defaultContentLanguageInSubdir = true
[languages]
[languages.en]
weight = 1
[languages.fr]
weight = 2
removePathAccents = true
-- content/διακριτικός.md --
-- content/διακριτικός.fr.md --
-- layouts/_default/single.html --
{{ .Language.Lang }}|Single.
-- layouts/_default/list.html --
List
`
	b := hugolib.Test(t, files)

	b.AssertFileContent("public/en/διακριτικός/index.html", "en|Single")
	b.AssertFileContent("public/fr/διακριτικος/index.html", "fr|Single")
}

func TestDisablePathToLower(t *testing.T) {
	t.Parallel()

	files := `
-- hugo.toml --
disableKinds = ["taxonomy", "term"]
defaultContentLanguage = "en"
defaultContentLanguageInSubdir = true
[languages]
[languages.en]
weight = 1
[languages.fr]
weight = 2
disablePathToLower = true
-- content/MySection/MyPage.md --
-- content/MySection/MyPage.fr.md --
-- content/MySection/MyBundle/index.md --
-- content/MySection/MyBundle/index.fr.md --
-- layouts/_default/single.html --
{{ .Language.Lang }}|Single.
-- layouts/_default/list.html --
{{ .Language.Lang }}|List.
`
	b := hugolib.Test(t, files)

	b.AssertFileContent("public/en/mysection/index.html", "en|List")
	b.AssertFileContent("public/en/mysection/mypage/index.html", "en|Single")
	b.AssertFileContent("public/fr/MySection/index.html", "fr|List")
	b.AssertFileContent("public/fr/MySection/MyPage/index.html", "fr|Single")
	b.AssertFileContent("public/en/mysection/mybundle/index.html", "en|Single")
	b.AssertFileContent("public/fr/MySection/MyBundle/index.html", "fr|Single")
}
