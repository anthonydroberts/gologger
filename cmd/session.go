package cmd

import (
	"fmt"

	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/spf13/cobra"
)

// sessionCmd represents the session command
var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Create, delete, view and switch between existing sessions",
	Long:  `Create, delete, view and switch between existing sessions (with no arguments this will display the active session)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		terminal.Msg("print", fmt.Sprintf("Current session: %s", state.Glog.ActiveSession))
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)
}
