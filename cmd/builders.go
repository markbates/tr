package cmd

type builderFunc func([]string) *Cmd
type builder struct {
	keyFile     string
	builderFunc builderFunc
}

var builders []builder

func init() {
	builders = []builder{
		{"Makefile", MakefileBuilder},
		{"./test.sh", TestSHBuilder},
		{".buffalo.dev.yml", BuffaloBuilder},
		{"**/*_test.go", GoBuilder},
		{"bin/rails", RailsBuilder},
		{"Rakefile", RakefileBuilder},
		{"package.json", NPMBuilder},
	}
}
