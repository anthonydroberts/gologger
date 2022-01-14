package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/anthonydroberts/gologger/data"
	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history [Number]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Browse & open previously saved command logs in the session",
	Long: "Browse & open previously saved command logs in the session\n" +
		"Example usage:\n" +
		"gologger history 0 --editor nano    --- opens the most recently saved log, in nano\n" +
		"gologger history                    --- opens an interactive list of saved logs to choose from\n",
	Run: func(cmd *cobra.Command, args []string) {
		entryContentList := data.GetEntries()
		var selectedEntry *data.Entry
		var entryListDisplay []string

		if len(entryContentList) == 0 {
			terminal.Msg("print", fmt.Sprintf("There are no existing entries in session '%s'. Create one with 'gologger run <Command>'", state.Glog.ActiveSession))
			os.Exit(0)
		}

		if len(args) == 0 {
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
				return
			}

			selectedEntry = entryContentList[numChoice]
		} else {
			inputNum, convertArgErr := strconv.Atoi(args[0])
			if convertArgErr != nil || inputNum+1 > len(entryContentList) {
				log.Fatalf("Argument must be a number between 0 and %d\n", len(entryContentList))
			}

			selectedEntry = entryContentList[inputNum]
		}

		terminal.Msg("success", fmt.Sprintf("Opening entry: '%s'", selectedEntry.Command))
		data.OpenEntryData(selectedEntry, cmd.Flag("editor").Value.String())
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().StringP("editor", "e", "terminal-output", "Open the log with a provided editor program name (vim, code, notepad, etc)")
}
