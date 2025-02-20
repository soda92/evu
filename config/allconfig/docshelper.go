package allconfig

import (
	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/docshelper"
)

// This is is just some helpers used to create some JSON used in the Hugo docs.
func init() {
	docsProvider := func() docshelper.DocProvider {
		cfg := config.New()
		for configRoot, v := range allDecoderSetups {
			if v.internalOrDeprecated {
				continue
			}
			cfg.Set(configRoot, make(maps.Params))
		}
		lang := maps.Params{
			"en": maps.Params{
				"menus":  maps.Params{},
				"params": maps.Params{},
			},
		}
		cfg.Set("languages", lang)
		cfg.SetDefaultMergeStrategy()

		configHelpers := map[string]any{
			"mergeStrategy": cfg.Get(""),
		}
		return docshelper.DocProvider{"config_helpers": configHelpers}
	}

	docshelper.AddDocProviderFunc(docsProvider)
}
