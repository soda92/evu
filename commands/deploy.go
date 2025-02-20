//go:build withdeploy

package commands

import (
	"context"

	"github.com/gohugoio/hugo/deploy"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
)

func newDeployCommand() simplecobra.Commander {
	return &simpleCommand{
		name:  "deploy",
		short: "Deploy your site to a cloud provider",
		long: `Deploy your site to a cloud provider

See https://gohugo.io/hosting-and-deployment/hugo-deploy/ for detailed
documentation.
`,
		run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
			h, err := r.Hugo(flagsToCfgWithAdditionalConfigBase(cd, nil, "deployment"))
			if err != nil {
				return err
			}
			deployer, err := deploy.New(h.Configs.GetFirstLanguageConfig(), h.Log, h.PathSpec.PublishFs)
			if err != nil {
				return err
			}
			return deployer.Deploy(ctx)
		},
		withc: func(cmd *cobra.Command, r *rootCommand) {
			applyDeployFlags(cmd, r)
		},
	}
}
