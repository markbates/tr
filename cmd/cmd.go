package cmd

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

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

func Exit(err error) {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	} else {
		log.Fatal(err)
	}
}

func Run(cmd *Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		Exit(err)
	}
	os.Exit(0)
}

var Exists = func(path string) bool {
	m, err := zglob.Glob(path)
	return err == nil && len(m) > 0
}
