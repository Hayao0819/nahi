package flist

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type walkResult struct {
	path     string
	isListed bool
	err      error
}

var successExeclude = walkResult{path: "", isListed: false, err: nil}

func singleWalker(dir string, opt *options) func(path string, d fs.DirEntry) walkResult {

	return func(path string, d fs.DirEntry) walkResult {
		// ファイルのみでディレクトリは除外
		if opt.fileOnly && d.IsDir() {
			return successExeclude
		}

		// ディレクトリのみでファイルは除外
		if opt.dirOnly && !d.IsDir() {
			return successExeclude
		}

		// 階層の深さ
		depthBasis := len(strings.Split(dir, string(os.PathSeparator)))
		depth := len(strings.Split(path, string(os.PathSeparator)))
		depthDiff := depth - depthBasis
		if !(opt.maxDepth < 0) && depthDiff > opt.maxDepth {
			return successExeclude
		}
		if depthDiff < opt.minDepth {
			return successExeclude
		}

		// 実行可能ファイルのみ
		if opt.executableOnly {
			info, err := os.Stat(path)
			if err != nil {
				return walkResult{path: "", isListed: false, err: err}
			}
			if !isExecutable(info.Mode()) {
				return successExeclude
			}
		}

		// 拡張子
		if len(opt.extsOnly) > 0 {
			ext := filepath.Ext(path)
			if !slices.Contains(opt.extsOnly, ext) {
				return successExeclude
			}
		}

		// 出力形式
		rtnPath := path
		if opt.relpath {
			rtnPath = strings.Replace(rtnPath, dir, ".", 1)
		}
		if opt.filename {
			rtnPath = filepath.Base(rtnPath)
		}

		// return true, rtnPath, nil
		return walkResult{path: rtnPath, isListed: true, err: nil}
	}

}
