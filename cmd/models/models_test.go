package models

import (
	"io/ioutil"

	"github.com/markbates/going/randx"
)

func tx(fn func()) {
	od := dbLocation
	defer func() {
		dbLocation = od
	}()
	dbLocation = func() string {
		dir, _ := ioutil.TempDir("", "tt")
		path := dir + randx.String(20) + ".db"
		return path
	}

	err := connect()
	if err != nil {
		panic(err)
	}

	fn()
}
