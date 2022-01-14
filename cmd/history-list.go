package cmd

import (
	"fmt"
	"os"

	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var historyListCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(0),
	Short: "Prints a formatted table of all saved logs in the session",
	Long:  `Prints a formatted table of all saved logs in the session`,
	Run: func(cmd *cobra.Command, args []string) {
		entriesList := data.GetEntries()

		if len(entriesList) == 0 {
			terminal.Msg("print", fmt.Sprintf("There are no existing entries in session '%s'. Create one with 'gologger run <Command>'", state.Glog.ActiveSession))
			os.Exit(0)
		}

		terminal.Msg("print", fmt.Sprintf("History for active session: %s", state.Glog.ActiveSession))

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Entry Command", "Date", "Status / Exit Code"})
		for i, entry := range entriesList {
			entryCommand := entry.Command
			entryDate := entry.TimeStamp
			entryExitCode := entry.ExitCode

			exitCodeColor := color.New(color.FgGreen, color.Bold)
			if entryExitCode != 0 {
				exitCodeColor = color.New(color.FgRed, color.Bold)
			}

			t.AppendRow([]interface{}{i, entryCommand, entryDate, exitCodeColor.Sprintf("%d", entryExitCode)})
		}

		t.Render()
	},
}

func init() {
	historyCmd.AddCommand(historyListCmd)
}
