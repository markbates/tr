package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use: "make",
	Run: func(cmd *cobra.Command, args []string) {
		Run(RunMakefile(args))
	},
}

func RunMakefile(args []string) *Cmd {
	fmt.Println("Testing via Makefile")
	cmd := New("make", "test")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}
