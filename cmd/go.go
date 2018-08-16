package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
	cmd := New(goBin, "test")
	if strings.Contains(goBin, "vgo") {
		cmd.Args = append(cmd.Args, "-vet", "off")
	}
	cmd.Args = append(cmd.Args, args...)
	runFlag := false
	for _, a := range cmd.Args {
		if a == "-run" {
			runFlag = true
		}
	}
	if !runFlag {
		c := exec.Command(goBin, "list", "./...")
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
