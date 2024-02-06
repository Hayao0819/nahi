package filelist

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type fileListOptions = struct {
	recursive bool
	maxDepth  int
	minDepth  int
	fileOnly  bool
	dirOnly   bool
}

func WithRecursive() func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.recursive = true
	}
}

func WithMaxDepth(depth int) func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.maxDepth = depth
	}
}

func WithMinDepth(depth int) func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.minDepth = depth
	}
}

func WithFileOnly() func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.fileOnly = true
	}
}

func WithDirOnly() func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.dirOnly = true
	}
}

func isDirEntryShouldSkip(d fs.DirEntry, fileOnly, dirOnly bool) bool {
	if fileOnly && d.IsDir() {
		return true
	}
	if dirOnly && !d.IsDir() {
		return true
	}
	return false
}

func FileList(dir string, opts ...func(*fileListOptions)) (*[]string, error) {
	opt := &fileListOptions{
		recursive: false,
		maxDepth:  -1,
		minDepth:  0,
		fileOnly:  false,
		dirOnly:   false,
	}

	for _, o := range opts {
		o(opt)
	}

	rtn := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// handle error
		if err != nil {
			return err
		}

		// check min depth
		if d.IsDir() && (strings.Count(path, string(os.PathSeparator)) > opt.maxDepth || strings.Count(path, string(os.PathSeparator)) < opt.minDepth) {
			fmt.Println("skip", path)
			return fs.SkipDir
		}

		if isDirEntryShouldSkip(d, opt.fileOnly, opt.dirOnly) {
			return nil
		}

		rtn = append(rtn, path)

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
