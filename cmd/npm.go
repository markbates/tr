package cmd

import "github.com/spf13/cobra"

var npmCmd = &cobra.Command{
	Use: "npm",
	Run: func(cmd *cobra.Command, args []string) {
		Run(NPMBuilder(args))
	},
}

func NPMBuilder(args []string) *Cmd {
	cmd := New("npm", "test")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func init() {
	RootCmd.AddCommand(npmCmd)
}
