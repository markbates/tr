package cmd

import (
	"os"
	"regexp"

	"github.com/gobuffalo/envy"
	"github.com/spf13/cobra"
)

const vendorPattern = "/vendor/"

var vendorRegex *regexp.Regexp

func init() {
	RootCmd.AddCommand(goCmd)
	vendorRegex = regexp.MustCompile(vendorPattern)
}

var goCmd = &cobra.Command{
	Use:                "go",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		Run(GoBuilder(args))
	},
}

var goBin = envy.Get("GO_BIN", "go")

func GoBuilder(args []string) *Cmd {
	os.Setenv("GO_ENV", "test")
	cmd := New(goBin, "test", "-cover")
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
