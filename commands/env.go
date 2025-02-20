package commands

import (
	"context"
	"runtime"

	"github.com/bep/simplecobra"
	"github.com/gohugoio/hugo/common/hugo"
	"github.com/spf13/cobra"
)

func newEnvCommand() simplecobra.Commander {
	return &simpleCommand{
		name:  "env",
		short: "Display version and environment info",
		long:  "Display version and environment info. This is useful in Hugo bug reports",
		run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
			r.Printf("%s\n", hugo.BuildVersionString())
			r.Printf("GOOS=%q\n", runtime.GOOS)
			r.Printf("GOARCH=%q\n", runtime.GOARCH)
			r.Printf("GOVERSION=%q\n", runtime.Version())

			if r.isVerbose() {
				deps := hugo.GetDependencyList()
				for _, dep := range deps {
					r.Printf("%s\n", dep)
				}
			} else {
				// These are also included in the GetDependencyList above;
				// always print these as these are most likely the most useful to know about.
				deps := hugo.GetDependencyListNonGo()
				for _, dep := range deps {
					r.Printf("%s\n", dep)
				}
			}
			return nil
		},
		withc: func(cmd *cobra.Command, r *rootCommand) {
			cmd.ValidArgsFunction = cobra.NoFileCompletions
		},
	}
}

func newVersionCmd() simplecobra.Commander {
	return &simpleCommand{
		name: "version",
		run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
			r.Println(hugo.BuildVersionString())
			return nil
		},
		short: "Display version",
		long:  "Display version and environment info. This is useful in Hugo bug reports.",
		withc: func(cmd *cobra.Command, r *rootCommand) {
			cmd.ValidArgsFunction = cobra.NoFileCompletions
		},
	}
}
