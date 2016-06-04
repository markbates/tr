package cmd

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/markbates/going/clam"
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
		Run(RunGoTests(args))
	},
}

func RunGoTests(args []string) *Cmd {
	os.Setenv("GO_ENV", "test")
	cmd := New("go", "test")
	cmd.Args = append(cmd.Args, args...)
	runFlag := false
	for _, a := range cmd.Args {
		if a == "-run" {
			runFlag = true
		}
	}
	if !runFlag {
		res, err := clam.Run(exec.Command("go", "list", "./..."))
		if err != nil {
			Exit(err)
		}
		pkgs := strings.Split(res, "\n")
		for _, p := range pkgs {
			if !vendorRegex.Match([]byte(p)) {
				cmd.Args = append(cmd.Args, p)
			}
		}
	}
	return cmd
}
