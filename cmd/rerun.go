package cmd

import (
	"github.com/markbates/tt/cmd/models"
	"github.com/spf13/cobra"
)

// rerunCmd represents the rerun command
var rerunCmd = &cobra.Command{
	Use:     "rerun",
	Aliases: []string{"rr"},
	Run: func(cmd *cobra.Command, args []string) {
		h, err := models.LastHistory()
		if err != nil {
			Exit(err)
		}
		n := h.CmdArgs[0]
		ar := h.CmdArgs[1:]
		c := New(n, ar...)
		Run(c)
	},
}

func init() {
	RootCmd.AddCommand(rerunCmd)
}
