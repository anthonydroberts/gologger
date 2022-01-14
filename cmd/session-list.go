package cmd

import (
	"os"

	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// sessionListCmd represents the list command
var sessionListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a table with all existing sessions & information about them",
	Long:  `Print a table with all existing sessions & information about them`,
	Run: func(cmd *cobra.Command, args []string) {
		sessionList := data.GetSessions()

		if len(sessionList) == 0 {
			terminal.Msg("print", "There are no existing sessions. Create one with 'gologger session create <SessionName>'")
			os.Exit(0)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Session Name", "Last Modified Date", "Entry Count"})
		for i, session := range sessionList {
			sessionString := session
			sessionLMDate := data.GetSessionLastModifiedTime(session)
			sessionEntryCount := len(data.GetSessionEntryPaths(session))

			if session == state.Glog.ActiveSession {
				col := color.New(color.FgHiYellow, color.Bold)
				sessionString = col.Sprintf("%s (a)", session)
			}

			t.AppendRow([]interface{}{i, sessionString, sessionLMDate.Format("2006-01-02 15:04:05"), sessionEntryCount})
			t.AppendSeparator()
		}

		t.Render()
	},
}

func init() {
	sessionCmd.AddCommand(sessionListCmd)
}
