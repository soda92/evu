package esbuild

import (
	"testing"

	"github.com/gohugoio/hugo/media"

	"github.com/evanw/esbuild/pkg/api"

	qt "github.com/frankban/quicktest"
)

func TestToBuildOptions(t *testing.T) {
	c := qt.New(t)

	opts := Options{
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:         true,
		Target:         api.ESNext,
		Format:         api.FormatIIFE,
		Platform:       api.PlatformBrowser,
		SourcesContent: 1,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			Target:   "es2018",
			Format:   "cjs",
			Minify:   true,
			AvoidTDZ: true,
		},
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:            true,
		Target:            api.ES2018,
		Format:            api.FormatCommonJS,
		Platform:          api.PlatformBrowser,
		SourcesContent:    1,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			Target: "es2018", Format: "cjs", Minify: true,
			SourceMap: "inline",
		},
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:            true,
		Target:            api.ES2018,
		Format:            api.FormatCommonJS,
		Platform:          api.PlatformBrowser,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		SourcesContent:    1,
		Sourcemap:         api.SourceMapInline,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			Target: "es2018", Format: "cjs", Minify: true,
			SourceMap: "inline",
		},
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:            true,
		Target:            api.ES2018,
		Format:            api.FormatCommonJS,
		Platform:          api.PlatformBrowser,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		Sourcemap:         api.SourceMapInline,
		SourcesContent:    1,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			Target: "es2018", Format: "cjs", Minify: true,
			SourceMap: "external",
		},
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:            true,
		Target:            api.ES2018,
		Format:            api.FormatCommonJS,
		Platform:          api.PlatformBrowser,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		Sourcemap:         api.SourceMapExternal,
		SourcesContent:    1,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			JSX: "automatic", JSXImportSource: "preact",
		},
		InternalOptions: InternalOptions{
			MediaType: media.Builtin.JavascriptType,
			Stdin:     true,
		},
	}

	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:         true,
		Target:         api.ESNext,
		Format:         api.FormatIIFE,
		Platform:       api.PlatformBrowser,
		SourcesContent: 1,
		Stdin: &api.StdinOptions{
			Loader: api.LoaderJS,
		},
		JSX:             api.JSXAutomatic,
		JSXImportSource: "preact",
	})

	opts = Options{
		ExternalOptions: ExternalOptions{
			Drop: "console",
		},
	}
	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled.Drop, qt.Equals, api.DropConsole)
	opts = Options{
		ExternalOptions: ExternalOptions{
			Drop: "debugger",
		},
	}
	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled.Drop, qt.Equals, api.DropDebugger)

	opts = Options{
		ExternalOptions: ExternalOptions{
			Drop: "adsfadsf",
		},
	}
	c.Assert(opts.compile(), qt.ErrorMatches, `unsupported drop type: "adsfadsf"`)
}

func TestToBuildOptionsTarget(t *testing.T) {
	c := qt.New(t)

	for _, test := range []struct {
		target string
		expect api.Target
	}{
		{"es2015", api.ES2015},
		{"es2016", api.ES2016},
		{"es2017", api.ES2017},
		{"es2018", api.ES2018},
		{"es2019", api.ES2019},
		{"es2020", api.ES2020},
		{"es2021", api.ES2021},
		{"es2022", api.ES2022},
		{"es2023", api.ES2023},
		{"", api.ESNext},
		{"esnext", api.ESNext},
	} {
		c.Run(test.target, func(c *qt.C) {
			opts := Options{
				ExternalOptions: ExternalOptions{
					Target: test.target,
				},
				InternalOptions: InternalOptions{
					MediaType: media.Builtin.JavascriptType,
				},
			}

			c.Assert(opts.compile(), qt.IsNil)
			c.Assert(opts.compiled.Target, qt.Equals, test.expect)
		})
	}
}

func TestDecodeExternalOptions(t *testing.T) {
	c := qt.New(t)
	m := map[string]any{
		"platform": "node",
	}
	ext, err := DecodeExternalOptions(m)
	c.Assert(err, qt.IsNil)
	c.Assert(ext, qt.DeepEquals, ExternalOptions{
		SourcesContent: true,
		Platform:       "node",
	})

	opts := Options{
		ExternalOptions: ext,
	}
	c.Assert(opts.compile(), qt.IsNil)
	c.Assert(opts.compiled, qt.DeepEquals, api.BuildOptions{
		Bundle:         true,
		Target:         api.ESNext,
		Format:         api.FormatIIFE,
		Platform:       api.PlatformNode,
		SourcesContent: api.SourcesContentInclude,
	})
}
