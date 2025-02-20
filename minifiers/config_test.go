package minifiers_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/config/testconfig"
)

func TestConfig(t *testing.T) {
	c := qt.New(t)
	v := config.New()

	v.Set("minify", map[string]any{
		"disablexml": true,
		"tdewolff": map[string]any{
			"html": map[string]any{
				"keepwhitespace": false,
			},
		},
	})

	conf := testconfig.GetTestConfigs(nil, v).Base.Minify

	c.Assert(conf.MinifyOutput, qt.Equals, false)

	// explicitly set value
	c.Assert(conf.Tdewolff.HTML.KeepWhitespace, qt.Equals, false)
	// default value
	c.Assert(conf.Tdewolff.HTML.KeepEndTags, qt.Equals, true)
	c.Assert(conf.Tdewolff.CSS.KeepCSS2, qt.Equals, true)

	// `enable` flags
	c.Assert(conf.DisableHTML, qt.Equals, false)
	c.Assert(conf.DisableXML, qt.Equals, true)
}

func TestConfigLegacy(t *testing.T) {
	c := qt.New(t)
	v := config.New()

	// This was a bool < Hugo v0.58.
	v.Set("minify", true)

	conf := testconfig.GetTestConfigs(nil, v).Base.Minify
	c.Assert(conf.MinifyOutput, qt.Equals, true)
}

func TestConfigNewCommentOptions(t *testing.T) {
	c := qt.New(t)
	v := config.New()

	// setting the old options should automatically set the new options
	v.Set("minify", map[string]any{
		"tdewolff": map[string]any{
			"html": map[string]any{
				"keepConditionalComments": false,
			},
			"svg": map[string]any{
				"decimal": "5",
			},
		},
	})

	conf := testconfig.GetTestConfigs(nil, v).Base.Minify

	c.Assert(conf.Tdewolff.HTML.KeepSpecialComments, qt.Equals, false)
	c.Assert(conf.Tdewolff.SVG.Precision, qt.Equals, 5)

	// the new values should win, regardless of the contents of the old values
	v = config.New()
	v.Set("minify", map[string]any{
		"tdewolff": map[string]any{
			"html": map[string]any{
				"keepConditionalComments": false,
				"keepSpecialComments":     true,
			},
			"svg": map[string]any{
				"decimal":   "5",
				"precision": "10",
			},
		},
	})

	conf = testconfig.GetTestConfigs(nil, v).Base.Minify

	c.Assert(conf.Tdewolff.HTML.KeepSpecialComments, qt.Equals, true)
	c.Assert(conf.Tdewolff.SVG.Precision, qt.Equals, 10)
}
