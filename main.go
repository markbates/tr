package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/markbates/going/clam"
	"github.com/mattn/go-zglob"
)

type Cmd struct {
	*exec.Cmd
}

func (c Cmd) String() string {
	return strings.Join(c.Args, " ")
}

func New(name string, args ...string) *Cmd {
	return &Cmd{exec.Command(name, args...)}
}

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

func RunTestSH(args []string) *Cmd {
	fmt.Println("Testing via ./test.sh")
	return New("./test.sh", args...)
}

func RunRakefile(args []string) *Cmd {
	if Exists("Gemfile") {
		return RunBundler(args)
	}
	fmt.Println("Testing via Rakefile")
	return New("rake", args...)
}

func RunBundler(args []string) *Cmd {
	fmt.Println("Testing via Rakefile (bundler)")
	cmd := New(os.Getenv("GEM_HOME")+"/bin/bundle", "exec", "rake")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func RunGoTests(args []string) *Cmd {
	os.Setenv("GO_ENV", "test")
	cmd := New("go", "test")
	cmd.Args = append(cmd.Args, args...)
	runFlag := false
	for _, a := range cmd.Args {
		if a == "-run" {
			runFlag = true
		}
	}
	if !runFlag {
		if Exists("glide.lock") {
			fmt.Println("Testing via go test (glide)")
			res, err := clam.Run(exec.Command("glide", "novendor"))
			if err != nil {
				exit(err)
			}
			pkgs := strings.Split(res, "\n")
			for _, p := range pkgs {
				cmd.Args = append(cmd.Args, p)
			}
		} else {
			fmt.Println("Testing via go test")
			cmd.Args = append(cmd.Args, "./...")
		}
	}
	return cmd
}

func RunMakefile(args []string) *Cmd {
	fmt.Println("Testing via Makefile")
	cmd := New("make", "test")
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func main() {
	fmt.Println("Test Runner: v1.1.0")
	args := os.Args[1:]
	for path, runner := range runners {
		if Exists(path) {
			run(runner(args))
		}
	}
}

func exit(err error) {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	} else {
		log.Fatal(err)
	}
}

func run(cmd *Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		exit(err)
	}
	os.Exit(0)
}

var Exists = func(path string) bool {
	m, err := zglob.Glob(path)
	return err == nil && len(m) > 0
}
