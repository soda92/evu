package os_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

// Issue 9599
func TestReadDirWorkDir(t *testing.T) {
	t.Parallel()

	files := `
-- config.toml --
theme = "mytheme"
-- myproject.txt --
Hello project!
-- themes/mytheme/mytheme.txt --
Hello theme!
-- layouts/index.html --
{{ $entries := (readDir ".") }}
START:|{{ range $entry := $entries }}{{ if not $entry.IsDir }}{{ $entry.Name }}|{{ end }}{{ end }}:END:


  `

	b := hugolib.NewIntegrationTestBuilder(
		hugolib.IntegrationTestConfig{
			T:           t,
			TxtarString: files,
			NeedsOsFS:   true,
		},
	).Build()

	b.AssertFileContent("public/index.html", `
START:|config.toml|myproject.txt|:END:
`)
}

// Issue 9620
func TestReadFileNotExists(t *testing.T) {
	t.Parallel()

	files := `
-- config.toml --
-- layouts/index.html --
{{ $fi := (readFile "doesnotexist") }}
{{ if $fi }}Failed{{ else }}OK{{ end }}


  `

	b := hugolib.NewIntegrationTestBuilder(
		hugolib.IntegrationTestConfig{
			T:           t,
			TxtarString: files,
			NeedsOsFS:   true,
		},
	).Build()

	b.AssertFileContent("public/index.html", `
OK
`)
}
