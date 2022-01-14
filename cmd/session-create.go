package cmd

import (
	"fmt"
	"os"

	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/spf13/cobra"
)

// sessionCreateCmd represents the create command
var sessionCreateCmd = &cobra.Command{
	Use:   "create SessionName",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new session",
	Long:  `Create a new session`,
	Run: func(cmd *cobra.Command, args []string) {
		if data.SessionExists(args[0]) {
			terminal.Msg("fail", fmt.Sprintf("Session with name '%s' exists already", args[0]))
			os.Exit(1)
		}

		data.CreateSession(args[0])
		terminal.Msg("success", fmt.Sprintf("Session '%s' created", args[0]))

		if cmd.Flags().Changed("switch") {
			data.UpdateActiveSession(args[0])
			terminal.Msg("success", fmt.Sprintf("Switched to session '%s'", args[0]))
		}
	},
}

func init() {
	sessionCmd.AddCommand(sessionCreateCmd)
	sessionCreateCmd.Flags().BoolP("switch", "s", false, "Switch to the new session after creation")
}
