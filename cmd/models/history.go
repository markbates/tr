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
	ID       uint64    `json:"id"`
	Time     time.Time `json:"time"`
	CmdArgs  []string  `json:"cmd"`
	Results  string    `json:"results"`
	Error    string    `json:"error"`
	ExitCode int       `json:"exit_code"`
}

func (h History) Bytes() []byte {
	b, _ := json.Marshal(h)
	return b
}

func (h History) String() string {
	return strings.Join(h.CmdArgs, " ")
}

func (h History) TimeString() string {
	return h.Time.In(time.Local).Format(time.RubyDate)
}

func (h History) Verdict() string {
	if h.ExitCode == 0 && h.Error == "" {
		return "PASS"
	}
	return "FAIL"
}

func (h History) Print() {
	fmt.Println(h.TimeString())
	fmt.Println(h.String())
	fmt.Println(string(h.Results))
	if h.Error != "" {
		fmt.Println(h.Error)
	}
}

func (h History) PrintShort() {
	fmt.Printf("%d)\t%s\t| %s\n", h.ID, h.TimeString(), h.Verdict())
}

func (h History) PrintShortVerbose() {
	fmt.Printf("%d)\t%s\t| %s\n\t%s\n", h.ID, h.TimeString(), h.Verdict(), h.String())
}

func (h *History) Save() error {
	db, err := DB()
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
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
	db, err := DB()
	if err != nil {
		return histories, err
	}
	err = db.View(func(tx *bolt.Tx) error {
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
	histories := Histories{}
	db, err := DB()
	if err != nil {
		return histories, err
	}
	err = db.View(func(tx *bolt.Tx) error {
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
	db, err := DB()
	if err != nil {
		return h, err
	}
	err = db.View(func(tx *bolt.Tx) error {
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
