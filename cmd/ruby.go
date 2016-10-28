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
	if len(args) == 0 {
		fmt.Println("You must supply a file name.")
		os.Exit(1)
	}

	var tf string

	f := args[0]
	rx := regexp.MustCompile("^(test|spec)")
	if rx.MatchString(f) {
		tf = f
		// cmd.Args = append(cmd.Args, f)
		// return cmd
	}

	if tf == "" {
		f_test := strings.Replace(f, ".rb", "_test.rb", 1)
		f_spec := strings.Replace(f, ".rb", "_spec.rb", 1)

		files := []string{
			f_test,
			f_spec,
			"test/" + strings.Replace(f_test, "app/", "", 1),
			"spec/" + strings.Replace(f_spec, "app/", "", 1),
		}

		for _, f := range files {
			if Exists(f) {
				tf = f
				break
				// cmd.Args = append(cmd.Args, f)
				// return cmd
			}
		}
		if tf == "" {
			fmt.Printf("Could not find a corresponding test for %s\n", f)
			os.Exit(1)
		}
	}

	rx = regexp.MustCompile("spec.rb")
	if rx.MatchString(tf) {
		cmd := BundlerBuilder("rspec")
		cmd.Args = append(cmd.Args, tf)
		return cmd
	}

	cmd := BundlerBuilder("ruby")
	cmd.Args = append(cmd.Args, "-Itest", tf)
	return cmd
}

func init() {
	RootCmd.AddCommand(rubyCmd)
}
