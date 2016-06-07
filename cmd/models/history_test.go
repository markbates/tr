package models

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_History_Verdict(t *testing.T) {
	r := require.New(t)
	h := History{}
	r.Equal("PASS", h.Verdict())
	h.Error = "some error"
	h.ExitCode = 1
	r.Equal("FAIL", h.Verdict())
}

func Test_History_Save(t *testing.T) {
	r := require.New(t)
	h := &History{
		CmdArgs:  []string{"foo", "bar"},
		Results:  []byte("some results"),
		Error:    "some error",
		ExitCode: 500,
	}
	tx(func() {
		err := h.Save()
		r.NoError(err)

		r.NotZero(h.ID)
		r.NotZero(h.Time)

		s := strconv.FormatUint(h.ID, 10)
		hs, err := GetHistories([]string{s})
		r.Equal(1, len(hs))
		r.NoError(err)
	})
}

func Test_AllHistories(t *testing.T) {
	r := require.New(t)
	h := History{
		CmdArgs:  []string{"foo", "bar"},
		Results:  []byte("some results"),
		Error:    "some error",
		ExitCode: 500,
	}
	hs := Histories{h, h, h}
	tx(func() {
		for _, h := range hs {
			err := h.Save()
			r.NoError(err)
		}

		hss, err := AllHistories()
		r.NoError(err)
		r.Equal(len(hs), len(hss))
	})
}

func Test_LastHistory(t *testing.T) {
	r := require.New(t)
	h := History{
		CmdArgs:  []string{"foo", "bar"},
		Results:  []byte("some results"),
		Error:    "some error",
		ExitCode: 500,
	}
	hs := Histories{h, h, h}
	tx(func() {
		for _, h := range hs {
			err := h.Save()
			r.NoError(err)
		}

		h, err := LastHistory()
		r.NoError(err)
		r.Equal(uint64(3), h.ID)
	})
}
