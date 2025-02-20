package commands

import (
	"context"

	"github.com/bep/simplecobra"
)

// newExec wires up all of Hugo's CLI.
func newExec() (*simplecobra.Exec, error) {
	rootCmd := &rootCommand{
		commands: []simplecobra.Commander{
			newHugoBuildCmd(),
			newVersionCmd(),
			newEnvCommand(),
			newServerCommand(),
			newDeployCommand(),
			newConfigCommand(),
			newNewCommand(),
			newConvertCommand(),
			newImportCommand(),
			newListCommand(),
			newModCommands(),
			newGenCommand(),
			newReleaseCommand(),
		},
	}

	return simplecobra.New(rootCmd)
}

func newHugoBuildCmd() simplecobra.Commander {
	return &hugoBuildCommand{}
}

// hugoBuildCommand just delegates to the rootCommand.
type hugoBuildCommand struct {
	rootCmd *rootCommand
}

func (c *hugoBuildCommand) Commands() []simplecobra.Commander {
	return nil
}

func (c *hugoBuildCommand) Name() string {
	return "build"
}

func (c *hugoBuildCommand) Init(cd *simplecobra.Commandeer) error {
	c.rootCmd = cd.Root.Command.(*rootCommand)
	return c.rootCmd.initRootCommand("build", cd)
}

func (c *hugoBuildCommand) PreRun(cd, runner *simplecobra.Commandeer) error {
	return c.rootCmd.PreRun(cd, runner)
}

func (c *hugoBuildCommand) Run(ctx context.Context, cd *simplecobra.Commandeer, args []string) error {
	return c.rootCmd.Run(ctx, cd, args)
}
