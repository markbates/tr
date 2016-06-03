package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/markbates/going/clam"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use: "go",
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
		if Exists("glide.lock") {
			fmt.Println("Testing via go test (glide)")
			res, err := clam.Run(exec.Command("glide", "novendor"))
			if err != nil {
				Exit(err)
			}
			pkgs := strings.Split(res, "\n")
			for _, p := range pkgs {
				cmd.Args = append(cmd.Args, p)
			}
		} else {
			fmt.Println("Testing via go test")
			cmd.Args = append(cmd.Args, "./...")
		}
	}
	return cmd
}
