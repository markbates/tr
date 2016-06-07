package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(shCmd)
}

var shCmd = &cobra.Command{
	Use: "sh",
	Run: func(cmd *cobra.Command, args []string) {
		Run(TestSHBuilder(args))
	},
}

func TestSHBuilder(args []string) *Cmd {
	return New("./test.sh", args...)
}
