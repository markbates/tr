package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(shCmd)
}

var shCmd = &cobra.Command{
	Use: "sh",
	Run: func(cmd *cobra.Command, args []string) {
		Run(RunTestSH(args))
	},
}

func RunTestSH(args []string) *Cmd {
	fmt.Println("Testing via ./test.sh")
	return New("./test.sh", args...)
}
