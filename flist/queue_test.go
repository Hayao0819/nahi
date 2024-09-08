package flist_test

import (
	"os"
	"testing"

	"github.com/Hayao0819/nahi/flist"
)

func handleError(t *testing.T, errs ...error) {
	for _, err := range errs {
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
}

func TestQueueFileList(t *testing.T) {
	home, err := os.Getwd()
	handleError(t, err)

	testcases := [][]flist.Queue{
		{
			*flist.NewQueue("rel-max1", home, flist.WithMaxDepth(1), flist.WithRelPath()),
			*flist.NewQueue("abs-max3", home, flist.WithMaxDepth(3)),
		},
	}

	for _, tc := range testcases {
		list, errs := flist.GetAll(tc...)
		for i, l := range list {
			t.Logf("list-%s: %v", i, *l)
		}
		if len(errs) > 0 {
			handleError(t, err)
		}
	}
}
