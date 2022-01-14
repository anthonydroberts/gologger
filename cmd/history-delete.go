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

// historyDeleteCmd represents the delete command
var historyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(0),
	Short: "Browse & delete previously saved command logs in the session",
	Long:  `Browse & open previously saved command logs in the session`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("all") {
			terminal.Msg("print", fmt.Sprintf("Deleting all entries in session: '%s'", state.Glog.ActiveSession))
			data.DeleteEntries(state.Glog.ActiveSession)
			os.Exit(0)
		}

		entryContentList := data.GetEntries()
		var selectedEntry *data.Entry
		var entryListDisplay []string

		if len(entryContentList) == 0 {
			terminal.Msg("print", fmt.Sprintf("There are no existing entries in session '%s'. Create one with 'gologger run <Command>'", state.Glog.ActiveSession))
			os.Exit(0)
		}

		for _, entry := range entryContentList {
			entryListDisplay = append(entryListDisplay, fmt.Sprintf("%s | Status: %d > %s", entry.TimeStamp, entry.ExitCode, entry.Command))
		}

		prompt := promptui.Select{
			Label: fmt.Sprintf("Select an entry to view (current session: '%s')", state.Glog.ActiveSession),
			Items: entryListDisplay,
		}
		numChoice, _, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed or cancelled %v\n", err)
		}

		selectedEntry = entryContentList[numChoice]

		terminal.Msg("print", fmt.Sprintf("Deleting entry: '%s'", selectedEntry.Command))
		data.DeleteEntry(selectedEntry)
	},
}

func init() {
	historyCmd.AddCommand(historyDeleteCmd)
	historyDeleteCmd.Flags().BoolP("all", "a", false, "Delete all existing entries in the active session")
}
