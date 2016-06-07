package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/boltdb/bolt"
	"github.com/markbates/tt/cmd/models"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(historyCmd)
}

var historyCmd = &cobra.Command{
	Use: "history",
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
	err := os.Remove(location())
	if err != nil {
		Exit(err)
	}
}

func printHistory(v []byte) error {
	h := &models.History{}
	err := json.Unmarshal(v, h)
	if err != nil {
		return err
	}
	h.Print()
	if h.ExitCode != 0 {
		os.Exit(h.ExitCode)
	}
	return err
}

func showLastHistory() {
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("history"))
		c := b.Cursor()
		_, v := c.Last()
		if v != nil {
			return printHistory(v)
		}

		return nil
	})
}

func showHistories(args []string) {
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("history"))

		for _, ind := range args {
			printHistory(b.Get([]byte(ind)))
		}

		return nil
	})

}

func listHistories() {
	histories := models.Histories{}
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("history"))

		b.ForEach(func(k, v []byte) error {
			h := models.History{}
			err := json.Unmarshal(v, &h)
			if err != nil {
				return err
			}
			histories = append(histories, h)
			return nil
		})
		return nil
	})

	sort.Sort(histories)
	for k, h := range histories {
		fmt.Printf("%d)\t%s\t| %s\n\t%s\n", k+1, h.Time.In(time.Local), h.Verdict(), h.String())
	}
}
