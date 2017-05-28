package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(railsCmd)
}

var railsCmd = &cobra.Command{
	Use: "rails",
	Run: func(cmd *cobra.Command, args []string) {
		Run(RailsBuilder(args))
	},
}

func RailsBuilder(args []string) *Cmd {
	args = append([]string{"test"}, args...)
	return New("bin/rails", args...)
}
