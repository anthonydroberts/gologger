package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// sessionDeleteCmd represents the delete command
var sessionDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.MaximumNArgs(1),
	Short: "A brief description of your command",
	Long:  `sdelete long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		sessionList := data.GetSessions()
		var selectedSession string

		// Remove active session from possible selections
		for i, s := range sessionList {
			if s == state.Glog.ActiveSession {
				sessionList = append(sessionList[:i], sessionList[i+1:]...)
				break
			}
		}

		if len(sessionList) == 0 {
			terminal.Msg("print", "There are no existing non-active sessions. Create one with 'gologger session create <SessionName>'")
			os.Exit(0)
		}

		if len(args) == 0 {
			prompt := promptui.Select{
				Label: fmt.Sprintf("Select a session to delete (not including active session: %s)", state.Glog.ActiveSession),
				Items: sessionList,
			}
			_, choice, err := prompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed or cancelled %v\n", err)
				return
			}

			selectedSession = choice
		} else {
			selectedSession = args[0]
		}

		if selectedSession == state.Glog.ActiveSession {
			terminal.Msg("fail", fmt.Sprintf("Cannot delete '%s' as it is the active session", selectedSession))
			os.Exit(1)
		}

		if !data.SessionExists(selectedSession) {
			terminal.Msg("fail", fmt.Sprintf("Session '%s' does not exist", selectedSession))
			os.Exit(1)
		}

		data.DeleteSession(selectedSession)
		terminal.Msg("success", fmt.Sprintf("Deleted session '%s'", selectedSession))
	},
}

func init() {
	sessionCmd.AddCommand(sessionDeleteCmd)
}
