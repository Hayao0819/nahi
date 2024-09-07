package flist_test

import (
	"os"
	"testing"

	"github.com/Hayao0819/nahi/flist"
)

func handleError(t *testing.T, err ...error) {
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestFileList(t *testing.T) {
	home, err := os.Getwd()
	handleError(t, err)

	testcases := [][]flist.Queue{
		{
			*flist.NewQueue(home, flist.WithMaxDepth(1), flist.WithRelPath()),
			*flist.NewQueue(home, flist.WithMaxDepth(3)),
		},
	}

	for _, tc := range testcases {
		list, errs := flist.GetAll(tc...)
		t.Logf("list: %v", list)
		if len(errs) > 0 {
			handleError(t, err)
		}
	}
}
