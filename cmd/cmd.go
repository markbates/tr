package cmd

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/boltdb/bolt"
	"github.com/markbates/tt/cmd/models"
	"github.com/mattn/go-zglob"
	"github.com/mitchellh/go-homedir"
)

var DB *bolt.DB
var PWD string

func init() {
	var err error
	PWD, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DB, err = bolt.Open(location(), 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("history"))
		return err
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func Run(cmd *Cmd) {
	h := &models.History{
		CmdArgs: cmd.Args,
		Results: []byte{},
		Time:    time.Now(),
	}
	var err error
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("history"))
		fmt.Println(cmd.String())
		fmt.Println("------------------------------------\n")

		var bb bytes.Buffer
		w := io.MultiWriter(os.Stdout, &bb)

		cmd.Stdout = w
		cmd.Stderr = os.Stderr
		cmd.Start()
		err = cmd.Wait()
		if err != nil {
			h.Error = err.Error()
		}

		h.Results = bb.Bytes()
		h.ExitCode = exitStatus(err)
		id, _ := b.NextSequence()
		h.ID = id
		b.Put(itob(id), h.Bytes())

		return nil
	})
	if err != nil {
		Exit(err)
	}
	os.Exit(h.ExitCode)
}

var Exists = func(path string) bool {
	m, err := zglob.Glob(path)
	return err == nil && len(m) > 0
}

func location() string {
	dir, _ := homedir.Dir()
	dir, _ = homedir.Expand(dir)
	dir = fmt.Sprintf("%s/.tt", dir)
	os.MkdirAll(dir, 0755)
	l := fmt.Sprintf("%s/%s.db", dir, GetMD5Hash(PWD))
	return l
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func itob(v uint64) []byte {
	s := strconv.FormatUint(v, 10)
	return []byte(s)
}
