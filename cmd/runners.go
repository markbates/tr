package cmd

type runFunc func([]string) *Cmd

var runners map[string]runFunc

func init() {
	runners = map[string]runFunc{
		"./test.sh":    RunTestSH,
		"Makefile":     RunMakefile,
		"Rakefile":     RunRakefile,
		"**/*_test.go": RunGoTests,
	}
}
