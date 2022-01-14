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

// sessionSwitchCmd represents the switch command
var sessionSwitchCmd = &cobra.Command{
	Use:   "switch",
	Args:  cobra.MaximumNArgs(1),
	Short: "Switch between existing sessions",
	Long:  `Switch between existing sessions`,
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
				Label: fmt.Sprintf("Select session (current: '%s')", state.Glog.ActiveSession),
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

		if !data.SessionExists(selectedSession) {
			terminal.Msg("fail", fmt.Sprintf("Session '%s' does not exist", args[0]))
			os.Exit(1)
		}

		data.UpdateActiveSession(selectedSession)
		terminal.Msg("success", fmt.Sprintf("Switched to session '%s'", selectedSession))
	},
}

func init() {
	sessionCmd.AddCommand(sessionSwitchCmd)
}
