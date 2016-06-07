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
		Run(RakefileBuilder(args))
	},
}

func RakefileBuilder(args []string) *Cmd {
	cmd := BundlerBuilder("rake")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func BundlerBuilder(name string) *Cmd {
	if Exists("Gemfile") {
		return New(os.Getenv("GEM_HOME")+"/bin/bundle", "exec", name)
	}
	return New(name)
}
