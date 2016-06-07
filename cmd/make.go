package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use: "make",
	Run: func(cmd *cobra.Command, args []string) {
		Run(MakefileBuilder(args))
	},
}

func MakefileBuilder(args []string) *Cmd {
	cmd := New("make", "test")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}
