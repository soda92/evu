package source_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/config/testconfig"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/source"
	"github.com/spf13/afero"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/hugofs"
)

func TestIgnoreDotFilesAndDirectories(t *testing.T) {
	c := qt.New(t)

	tests := []struct {
		path                string
		ignore              bool
		ignoreFilesRegexpes any
	}{
		{".foobar/", true, nil},
		{"foobar/.barfoo/", true, nil},
		{"barfoo.md", false, nil},
		{"foobar/barfoo.md", false, nil},
		{"foobar/.barfoo.md", true, nil},
		{".barfoo.md", true, nil},
		{".md", true, nil},
		{"foobar/barfoo.md~", true, nil},
		{".foobar/barfoo.md~", true, nil},
		{"foobar~/barfoo.md", false, nil},
		{"foobar/bar~foo.md", false, nil},
		{"foobar/foo.md", true, []string{"\\.md$", "\\.boo$"}},
		{"foobar/foo.html", false, []string{"\\.md$", "\\.boo$"}},
		{"foobar/foo.md", true, []string{"foo.md$"}},
		{"foobar/foo.md", true, []string{".*", "\\.md$", "\\.boo$"}},
		{"foobar/.#content.md", true, []string{"/\\.#"}},
		{".#foobar.md", true, []string{"^\\.#"}},
	}

	for i, test := range tests {
		test := test
		c.Run(fmt.Sprintf("[%d] %s", i, test.path), func(c *qt.C) {
			c.Parallel()
			v := config.New()
			v.Set("ignoreFiles", test.ignoreFilesRegexpes)
			v.Set("publishDir", "public")
			afs := afero.NewMemMapFs()
			conf := testconfig.GetTestConfig(afs, v)
			fs := hugofs.NewFromOld(afs, v)
			ps, err := helpers.NewPathSpec(fs, conf, nil)
			c.Assert(err, qt.IsNil)

			s := source.NewSourceSpec(ps, nil, fs.Source)

			if ignored := s.IgnoreFile(filepath.FromSlash(test.path)); test.ignore != ignored {
				t.Errorf("[%d] File not ignored", i)
			}
		})

	}
}
