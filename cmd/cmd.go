package cmd

import (
	"os/exec"
	"strings"
)

type Cmd struct {
	*exec.Cmd
}

func (c Cmd) String() string {
	return strings.TrimSpace(strings.Join(c.Args, " "))
}

func New(name string, args ...string) *Cmd {
	return &Cmd{exec.Command(name, args...)}
}
