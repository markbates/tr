package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
)

var DB *bolt.DB
var PWD string

const BucketName = "history"

func init() {
	var err error
	PWD, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DB, err = bolt.Open(dbLocation(), 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		return err
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ClearDB() error {
	return os.Remove(dbLocation())
}

func hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func dbLocation() string {
	dir, _ := homedir.Dir()
	dir, _ = homedir.Expand(dir)
	dir = fmt.Sprintf("%s/.tt", dir)
	os.MkdirAll(dir, 0755)
	l := fmt.Sprintf("%s/%s.db", dir, hash(PWD))
	return l
}

func itob(v uint64) []byte {
	s := strconv.FormatUint(v, 10)
	return []byte(s)
}