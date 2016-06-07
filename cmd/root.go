package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var RootCmd = &cobra.Command{
	Use:                "tr",
	DisableFlagParsing: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Test Runner: v%s\n\n", Version)
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, r := range builders {
			if Exists(r.keyFile) {
				Run(r.builderFunc(args))
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
