package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type History struct {
	ID       uint64
	Time     time.Time
	CmdArgs  []string
	Results  []byte
	Error    string
	ExitCode int
}

func (h History) Bytes() []byte {
	b, _ := json.Marshal(h)
	return b
}

func (h History) String() string {
	return strings.Join(h.CmdArgs, " ")
}

func (h History) Verdict() string {
	if h.ExitCode == 0 && h.Error == "" {
		return "PASS"
	}
	return "FAIL"
}

func (h History) Print() {
	fmt.Println(h.Time.In(time.Local))
	fmt.Println(h.String())
	fmt.Println(string(h.Results))
	if h.Error != "" {
		fmt.Println(h.Error)
	}
}

func (h History) PrintShort() {
	fmt.Printf("%d)\t%s\t| %s\n\t%s\n", h.ID, h.Time.In(time.Local), h.Verdict(), h.String())
}

type Histories []History

func (a Histories) Len() int {
	return len(a)
}

func (a Histories) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Histories) Less(i, j int) bool {
	return a[i].Time.UnixNano() < a[j].Time.UnixNano()
}
