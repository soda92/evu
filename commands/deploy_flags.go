package commands

import (
	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/spf13/cobra"
)

func applyDeployFlags(cmd *cobra.Command, r *rootCommand) {
	cmd.ValidArgsFunction = cobra.NoFileCompletions
	cmd.Flags().String("target", "", "target deployment from deployments section in config file; defaults to the first one")
	_ = cmd.RegisterFlagCompletionFunc("target", cobra.NoFileCompletions)
	cmd.Flags().Bool("confirm", false, "ask for confirmation before making changes to the target")
	cmd.Flags().Bool("dryRun", false, "dry run")
	cmd.Flags().Bool("force", false, "force upload of all files")
	cmd.Flags().Bool("invalidateCDN", deployconfig.DefaultConfig.InvalidateCDN, "invalidate the CDN cache listed in the deployment target")
	cmd.Flags().Int("maxDeletes", deployconfig.DefaultConfig.MaxDeletes, "maximum # of files to delete, or -1 to disable")
	_ = cmd.RegisterFlagCompletionFunc("maxDeletes", cobra.NoFileCompletions)
	cmd.Flags().Int("workers", deployconfig.DefaultConfig.Workers, "number of workers to transfer files. defaults to 10")
	_ = cmd.RegisterFlagCompletionFunc("workers", cobra.NoFileCompletions)
}
