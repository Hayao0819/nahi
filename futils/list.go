package futils

import (
	"github.com/Hayao0819/nahi/flist"
)

func RecursionFileList(dir string) (*[]string, error) {
	return flist.Get(dir, flist.WithFileOnly())
}

func RecursionDirList(dir string) (*[]string, error) {
	return flist.Get(dir, flist.WithDirOnly())
}
