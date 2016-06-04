package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

type History struct {
	Time     time.Time
	CmdArgs  []string
	Results  []byte
	Error    string
	ExitCode int
}

func (h History) Bytes() []byte {
	b, _ := json.Marshal(h)
	return b
}

func (h History) String() string {
	return strings.Join(h.CmdArgs, " ")
}

func (h History) Verdict() string {
	if h.ExitCode == 0 && h.Error == "" {
		return "PASS"
	}
	return "FAIL"
}

func (h History) Print() {
	fmt.Println(h.Time.In(time.Local))
	fmt.Println(h.String())
	fmt.Println(string(h.Results))
	if h.Error != "" {
		fmt.Println(h.Error)
	}
}

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
	h := &History{}
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
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("history"))

		b.ForEach(func(k, v []byte) error {
			h := &History{}
			err := json.Unmarshal(v, h)
			if err != nil {
				return err
			}
			fmt.Printf("%s)\t%s | %s\n\t%s\n", k, h.Time.In(time.Local), h.Verdict(), h.String())
			return nil
		})
		return nil
	})
}
