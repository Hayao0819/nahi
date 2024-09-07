package flist_test

import (
	"os"
	"testing"

	"github.com/Hayao0819/nahi/flist"
)


func TestAllFileList(t *testing.T) {
	home, err := os.Getwd()
	handleError(t, err)

	testcases := []struct {
		options []flist.Option
	}{
		{
			options: []flist.Option{
				flist.WithMaxDepth(1),
				flist.WithRelPath(),
			},
		},
	}

	for _, tc := range testcases {
		list, err := flist.Get(home, tc.options...)
		t.Logf("list: %v", *list)
		handleError(t, err)
	}
}
