package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/markbates/tt/cmd/models"
	"github.com/mattn/go-zglob"
)

var Exists = func(path string) bool {
	m, err := zglob.Glob(path)
	return err == nil && len(m) > 0
}

type Cmd struct {
	*exec.Cmd
}

func (c Cmd) String() string {
	return strings.TrimSpace(strings.Join(c.Args, " "))
}

func New(name string, args ...string) *Cmd {
	return &Cmd{exec.Command(name, args...)}
}

func Exit(err error) {
	os.Exit(exitStatus(err))
}

func Run(cmd *Cmd) {
	h := &models.History{
		CmdArgs: cmd.Args,
		Results: []byte{},
		Time:    time.Now(),
	}
	fmt.Println(cmd.String())

	var bb bytes.Buffer
	w := io.MultiWriter(os.Stdout, &bb)

	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		h.Error = err.Error()
	}

	h.Results = bb.Bytes()
	h.ExitCode = exitStatus(err)
	err = h.Save()

	if err != nil {
		Exit(err)
	}
	os.Exit(h.ExitCode)
}

func exitStatus(err error) int {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
		return 1
	}
	return 0
}
