package cmd

import (
	"bytes"
	"io"
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
		bb := &bytes.Buffer{}
		mw := io.MultiWriter(bb, os.Stdout, os.Stderr)
		c.Stderr = mw
		c.Stdout = mw
		c.Stdin = os.Stdin

		if err := c.Run(); err != nil {
			Exit(err)
		}
		pkgs := bytes.Split(bytes.TrimSpace(bb.Bytes()), []byte("\n"))
		for _, p := range pkgs {
			if !vendorRegex.Match(p) {
				cmd.Args = append(cmd.Args, string(p))
			}
		}
	}
	return cmd
}
