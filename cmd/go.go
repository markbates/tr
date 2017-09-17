package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

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

func GoBuilder(args []string) *Cmd {
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
		c := exec.Command("go", "list", "./...")
		res, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(string(res))
			Exit(err)
		}
		pkgs := bytes.Split(bytes.TrimSpace(res), []byte("\n"))
		for _, p := range pkgs {
			if !vendorRegex.Match(p) {
				cmd.Args = append(cmd.Args, string(p))
			}
		}
	}
	return cmd
}
