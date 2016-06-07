package cmd

type builderFunc func([]string) *Cmd
type builder struct {
	keyFile     string
	builderFunc builderFunc
}

var builders []builder

func init() {
	builders = []builder{
		{"./test.sh", TestSHBuilder},
		{"Makefile", MakefileBuilder},
		{"Rakefile", RakefileBuilder},
		{"package.json", NPMBuilder},
		{"**/*_test.go", GoBuilder},
	}
}
