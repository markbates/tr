package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use:                "go",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		Run(GoBuilder(args))
	},
}

func GoBuilder(args []string) *Cmd {
	os.Setenv("GO_ENV", "test")
	cmd := New("go", "test", "-cover", "-short")
	if len(args) == 0 {
		args = append(args, "./...")
	}
	cmd.Args = append(cmd.Args, args...)
	for _, a := range cmd.Args {
		if a == "-run" {
			return cmd
		}
	}

	return cmd
}
