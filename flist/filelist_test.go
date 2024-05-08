package flist_test

import (
	"testing"

	"github.com/Hayao0819/nahi/flist"
)

func TestFileList(t *testing.T) {
	got, err := flist.Get(".")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Logf("FileList: %v", *got)
}
