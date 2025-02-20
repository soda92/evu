package hugolib

import "testing"

func TestConfigDir(t *testing.T) {
	t.Parallel()

	files := `
-- config/_default/params.toml --
a = "acp1"
d = "dcp1"
-- config/_default/config.toml --
[params]
a = "ac1"
b = "bc1"

-- hugo.toml --
baseURL = "https://example.com"
disableKinds = ["taxonomy", "term", "RSS", "sitemap", "robotsTXT", "page", "section"]
ignoreErrors = ["error-missing-instagram-accesstoken"]
[params]
a = "a1"
b = "b1"
c = "c1"
-- layouts/index.html --
Params: {{ site.Params}}
`
	b := Test(t, files)

	b.AssertFileContent("public/index.html", `
Params: map[a:acp1 b:bc1 c:c1 d:dcp1]


`)
}
