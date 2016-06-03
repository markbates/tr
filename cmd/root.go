package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "tr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Test Runner: v%s\n", Version)
		for path, runner := range runners {
			if Exists(path) {
				Run(runner(args))
			}
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
