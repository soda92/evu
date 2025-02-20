package httpcache_test

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/hugolib"
)

func TestConfigCustom(t *testing.T) {
	files := `
-- hugo.toml --
[httpcache]
[httpcache.cache.for]
includes = ["**gohugo.io**"]
[[httpcache.polls]]
low = "5s"
high = "32s"
[httpcache.polls.for]
includes = ["**gohugo.io**"]
		
	
`

	b := hugolib.Test(t, files)

	httpcacheConf := b.H.Configs.Base.HTTPCache
	compiled := b.H.Configs.Base.C.HTTPCache

	b.Assert(httpcacheConf.Cache.For.Includes, qt.DeepEquals, []string{"**gohugo.io**"})
	b.Assert(httpcacheConf.Cache.For.Excludes, qt.IsNil)

	pc := compiled.PollConfigFor("https://gohugo.io/foo.jpg")
	b.Assert(pc.Config.Low, qt.Equals, 5*time.Second)
	b.Assert(pc.Config.High, qt.Equals, 32*time.Second)
	b.Assert(compiled.PollConfigFor("https://example.com/foo.jpg").IsZero(), qt.IsTrue)
}

func TestConfigDefault(t *testing.T) {
	files := `
-- hugo.toml --
`
	b := hugolib.Test(t, files)

	compiled := b.H.Configs.Base.C.HTTPCache

	b.Assert(compiled.For("https://gohugo.io/posts.json"), qt.IsFalse)
	b.Assert(compiled.For("https://gohugo.io/foo.jpg"), qt.IsFalse)
	b.Assert(compiled.PollConfigFor("https://gohugo.io/foo.jpg").Config.Disable, qt.IsTrue)
}
