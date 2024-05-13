package flist

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type fileListOptions = struct {
	maxDepth int
	minDepth int
	fileOnly bool
	dirOnly  bool
	extOnly  string
	extsOnly []string
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

func WithExtOnly(ext string) func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.extOnly = ext
	}
}

func WithExtsOnly(exts []string) func(*fileListOptions) {
	return func(opt *fileListOptions) {
		opt.extsOnly = exts
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

func Get(dir string, opts ...func(*fileListOptions)) (*[]string, error) {
	opt := fileListOptions{
		maxDepth: -1,
		minDepth: 0,
		fileOnly: false,
		dirOnly:  false,
		extOnly:  "",
		extsOnly: []string{},
	}

	for _, o := range opts {
		o(&opt)
	}

	rtn := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// handle error
		if err != nil {
			return err
		}

		if opt.fileOnly && d.IsDir() {
			return nil
		}
		if opt.dirOnly && !d.IsDir() {
			return nil
		}

		depth := len(strings.Split(path, string(os.PathSeparator))) - 1
		if !(opt.maxDepth < 0) && depth > opt.maxDepth {
			return nil
		}
		if depth < opt.minDepth {
			return nil
		}

		if opt.extOnly != "" {
			ext := filepath.Ext(path)
			if ext != opt.extOnly && !slices.Contains(opt.extsOnly, ext) {
				return nil
			}
		}

		rtn = append(rtn, path)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
