package commands

import (
	"context"

	"github.com/bep/simplecobra"
	"github.com/gohugoio/hugo/releaser"
	"github.com/spf13/cobra"
)

// Note: This is a command only meant for internal use and must be run
// via "go run -tags release main.go release" on the actual code base that is in the release.
func newReleaseCommand() simplecobra.Commander {
	var (
		step     int
		skipPush bool
		try      bool
	)

	return &simpleCommand{
		name:  "release",
		short: "Release a new version of Hugo",
		run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
			rel, err := releaser.New(skipPush, try, step)
			if err != nil {
				return err
			}

			return rel.Run()
		},
		withc: func(cmd *cobra.Command, r *rootCommand) {
			cmd.Hidden = true
			cmd.ValidArgsFunction = cobra.NoFileCompletions
			cmd.PersistentFlags().BoolVarP(&skipPush, "skip-push", "", false, "skip pushing to remote")
			cmd.PersistentFlags().BoolVarP(&try, "try", "", false, "no changes")
			cmd.PersistentFlags().IntVarP(&step, "step", "", 0, "step to run (1: set new version 2: prepare next dev version)")
			_ = cmd.RegisterFlagCompletionFunc("step", cobra.FixedCompletions([]string{"1", "2"}, cobra.ShellCompDirectiveNoFileComp))
		},
	}
}
