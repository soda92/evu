package commands

import (
	// For time zone lookups on Windows without Go installed.
	// See #8892
	_ "time/tzdata"

	"github.com/spf13/cobra"
)

func init() {
	// This message to show to Windows users if Hugo is opened from explorer.exe
	cobra.MousetrapHelpText = `

  Hugo is a command-line tool for generating static websites.

  You need to open PowerShell and run Hugo from there.

  Visit https://gohugo.io/ for more information.`
}
