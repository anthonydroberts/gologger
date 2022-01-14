package cmd

import (
	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/process"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run Command",
	Args:  cobra.ExactArgs(1),
	Short: "Run a command & save the output to a session",
	Long:  `Run a command & save the output to a session (Ex. gologger run 'ls -a')`,
	Run: func(cmd *cobra.Command, args []string) {
		commandRan, resultOutput, resultError, resultExitCode, startTime := process.RunCommand(args[0], cmd.Flags().Changed("silent"), false)

		var resultString string = resultOutput
		if resultError != "" {
			resultString = resultString + "\n" + resultError
		}

		data.CreateEntry(commandRan, []byte(resultString), resultExitCode, startTime)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("silent", "s", false, "Hide the command's terminal output while running")
}
