package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rakeCmd)
}

var rakeCmd = &cobra.Command{
	Use: "rake",
	Run: func(cmd *cobra.Command, args []string) {
		Run(RunRakefile(args))
	},
}

func RunRakefile(args []string) *Cmd {
	if Exists("Gemfile") {
		return RunBundler(args)
	}
	return New("rake", args...)
}

func RunBundler(args []string) *Cmd {
	cmd := New(os.Getenv("GEM_HOME")+"/bin/bundle", "exec", "rake")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}
