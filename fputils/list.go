package fputils

import (
	"io/fs"
	"path/filepath"
)

func RecursionFileList(dir string) (*[]string, error) {
	files := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &files, nil
}

func RecursionDirList(dir string) (*[]string, error) {
	dirs := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &dirs, nil
}

func FileList(dir string) (*[]string, error) {
	files, err := filepath.Glob(dir)
	if err != nil {
		return nil, err
	}
	return &files, nil
}
