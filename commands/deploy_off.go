//go:build !withdeploy

package commands

import (
	"context"
	"errors"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
)

func newDeployCommand() simplecobra.Commander {
	return &simpleCommand{
		name: "deploy",
		run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
			return errors.New("deploy not supported in this version of Hugo; install a release with 'withdeploy' in the archive filename or build yourself with the 'withdeploy' build tag. Also see https://github.com/gohugoio/hugo/pull/12995")
		},
		withc: func(cmd *cobra.Command, r *rootCommand) {
			applyDeployFlags(cmd, r)
			cmd.Hidden = true
		},
	}
}
