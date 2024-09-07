package flist

import (
	"io/fs"
	"path/filepath"
)

var getDefault func() *options = func() *options {
	return &options{
		maxDepth:       -1,
		minDepth:       0,
		fileOnly:       false,
		dirOnly:        false,
		filename:       false,
		filenameMatch:  []string{},
		extsOnly:       []string{},
		relpath:        false,
		executableOnly: false,
	}
}

func Get(dir string, opts ...Option) (*[]string, error) {
	opt := getDefault()

	for _, o := range opts {
		o(opt)
	}

	rtn := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		wakler := singleWalker(dir, opt)
		res := wakler(path, d)
		if res.isListed {
			rtn = append(rtn, res.path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
