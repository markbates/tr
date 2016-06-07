package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/markbates/tt/cmd/models"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(historyCmd)
}

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"h"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			listHistories()
		} else {
			switch args[0] {
			case "clear":
				clearHistory()
			case "last":
				showLastHistory()
			default:
				showHistories(args)
			}
		}
	},
}

func clearHistory() {
	err := models.ClearDB()
	if err != nil {
		Exit(err)
	}
	fmt.Println("History has been cleared.")
}

func showLastHistory() {
	h, err := models.LastHistory()
	if err != nil {
		fmt.Println("This is no last piece of History.")
		os.Exit(0)
	}
	h.Print()
}

func showHistories(args []string) {
	histories, err := models.GetHistories(args)
	if err != nil {
		Exit(err)
	}

	if len(histories) == 0 {
		fmt.Printf("There is no history for %s.\n", strings.Join(args, ", "))
		return
	}

	for _, h := range histories {
		h.Print()
	}
}

func listHistories() {
	histories, err := models.AllHistories()
	if err != nil {
		Exit(err)
	}

	if len(histories) == 0 {
		fmt.Println("There is no history.")
		return
	}

	for _, h := range histories {
		h.PrintShort()
	}
}
