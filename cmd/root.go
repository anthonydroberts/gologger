package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "gologger [command]",
	Long: `Gologger is a cross-platform productivity tool that enables quick & easy logging for terminal commands`,
}

func Execute() {
	// Disable Cobra's completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	cobra.CheckErr(rootCmd.Execute())
}
