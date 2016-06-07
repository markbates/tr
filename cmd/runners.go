package cmd

type runFunc func([]string) *Cmd
type runner struct {
	keyFile string
	runFunc runFunc
}

var runners []runner

func init() {
	runners = []runner{
		{"./test.sh", RunTestSH},
		{"Makefile", RunMakefile},
		{"Rakefile", RunRakefile},
		{"package.json", RunTestNPM},
		{"**/*_test.go", RunGoTests},
	}
}
