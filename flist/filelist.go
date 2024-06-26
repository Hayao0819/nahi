package flist

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func Get(dir string, opts ...Option) (*[]string, error) {
	opt := options{
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

		// フィルター

		// ファイルのみでディレクトリは除外
		if opt.fileOnly && d.IsDir() {
			return nil
		}

		// ディレクトリのみでファイルは除外
		if opt.dirOnly && !d.IsDir() {
			return nil
		}

		// 階層の深さ
		depthBasis := len(strings.Split(dir, string(os.PathSeparator)))
		depth := len(strings.Split(path, string(os.PathSeparator)))
		depthDiff := depth - depthBasis
		if !(opt.maxDepth < 0) && depthDiff > opt.maxDepth {
			return nil
		}
		if depthDiff < opt.minDepth {
			return nil
		}

		// 実行可能ファイルのみ
		if opt.executableOnly {
			info, err := os.Stat(path)
			if err != nil {
				return nil
			}
			if !isExecutable(info.Mode()) {
				return nil
			}
		}

		// 拡張子
		if opt.extOnly != "" {
			ext := filepath.Ext(path)
			if ext != opt.extOnly && !slices.Contains(opt.extsOnly, ext) {
				return nil
			}
		}

		// 出力形式
		rtnPath := path
		if opt.relpath {
			rtnPath = strings.Replace(rtnPath, dir, "", 1)
		}
		if opt.filename {
			rtnPath = filepath.Base(rtnPath)
		}

		rtn = append(rtn, rtnPath)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
