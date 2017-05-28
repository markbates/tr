package cmd

type builderFunc func([]string) *Cmd
type builder struct {
	keyFile     string
	builderFunc builderFunc
}

var builders []builder

func init() {
	builders = []builder{
		{".buffalo.dev.yml", BuffaloBuilder},
		{"./test.sh", TestSHBuilder},
		{"Makefile", MakefileBuilder},
		{"bin/rails", RailsBuilder},
		{"Rakefile", RakefileBuilder},
		{"**/*_test.go", GoBuilder},
		{"package.json", NPMBuilder},
	}
}
