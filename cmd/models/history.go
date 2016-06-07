package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/boltdb/bolt"
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

func (h *History) Save() error {
	err := DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		h.Time = time.Now()
		h.ID = id
		err = b.Put(itob(id), h.Bytes())
		return err
	})
	return err
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

func AllHistories() (Histories, error) {
	histories := Histories{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))

		return b.ForEach(func(k, v []byte) error {
			h := History{}
			err := json.Unmarshal(v, &h)
			if err != nil {
				return err
			}
			histories = append(histories, h)
			return nil
		})
	})

	sort.Sort(histories)
	return histories, err
}

func GetHistories(args []string) (Histories, error) {
	fmt.Printf("### args -> %#v\n", args)
	histories := Histories{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))

		for _, ind := range args {
			h := History{}
			v := b.Get([]byte(ind))
			if v != nil {
				err := json.Unmarshal(v, &h)
				if err != nil {
					return err
				}
				histories = append(histories, h)
			}
		}

		return nil
	})
	sort.Sort(histories)
	return histories, err
}

func LastHistory() (History, error) {
	h := History{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		c := b.Cursor()
		_, v := c.Last()
		if v != nil {
			err := json.Unmarshal(v, &h)
			return err
		}
		return errors.New("No history could be found.")
	})
	return h, err
}
