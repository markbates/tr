package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// rubyCmd represents the ruby command
var rubyCmd = &cobra.Command{
	Use: "ruby",
	Run: func(cmd *cobra.Command, args []string) {
		Run(RubyBuilder(args))
	},
}

func RubyBuilder(args []string) *Cmd {
	cmd := BundlerBuilder("ruby")
	cmd.Args = append(cmd.Args, "-Itest", "-Ispec")

	if len(args) == 0 {
		fmt.Println("You must supply a file name.")
		os.Exit(1)
	}
	f := args[0]
	rx := regexp.MustCompile("^(test|spec)")
	if rx.Match([]byte(f)) {
		cmd.Args = append(cmd.Args, f)
		return cmd
	}
	f_test := strings.Replace(f, ".rb", "_test.rb", 1)
	f_spec := strings.Replace(f, ".rb", "_spec.rb", 1)

	files := []string{
		f_test,
		f_spec,
		"test/" + strings.Replace(f_test, "app/", "", 1),
		"spec/" + strings.Replace(f_spec, "app/", "", 1),
	}

	for _, f := range files {
		_, err := os.Stat(f)
		if err == nil {
			cmd.Args = append(cmd.Args, f)
			return cmd
		}
	}
	return cmd
}

func init() {
	RootCmd.AddCommand(rubyCmd)
}
