package filelist_test

import (
	"testing"

	"github.com/Hayao0819/nahi/filelist"
)

func TestFileList(t *testing.T) {
	got, err := filelist.FileList(".")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Logf("FileList: %v", *got)
}
