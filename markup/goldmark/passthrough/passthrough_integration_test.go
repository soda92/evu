package passthrough_test

import (
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestPassthroughRenderHook(t *testing.T) {
	t.Parallel()

	files := `
-- config.toml --
[markup.goldmark.extensions.passthrough]
enable = true
[markup.goldmark.extensions.passthrough.delimiters]
block = [['$$', '$$']]
inline = [['$', '$']]
-- content/p1.md --
---
title: "p1"
---
## LaTeX test

Some inline LaTeX 1: $a^*=x-b^*$.

Block equation that would be mangled by default parser:

$$a^*=x-b^*$$

Some inline LaTeX 2: $a^*=x-b^*$.

-- layouts/_default/single.html --
{{ .Content }}
-- layouts/_default/_markup/render-passthrough-block.html --
Passthrough block: {{ .Inner | safeHTML }}|{{ .Type }}|{{ .Ordinal }}:END
-- layouts/_default/_markup/render-passthrough-inline.html --
Passthrough inline: {{ .Inner | safeHTML }}|{{ .Type }}|{{ .Ordinal }}:END

`

	b := hugolib.Test(t, files)
	b.AssertFileContent("public/p1/index.html", `
		Some inline LaTeX 1: Passthrough inline: a^*=x-b^*|inline|0:END
		Passthrough block: a^*=x-b^*|block|1:END
		Some inline LaTeX 2: Passthrough inline: a^*=x-b^*|inline|2:END

	`)
}
