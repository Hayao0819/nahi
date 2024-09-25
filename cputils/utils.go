package cputils

import (
	"log/slog"
	"os"
	"path/filepath"
)

func OnlySpecificExtention(ext string) func(srcinfo os.FileInfo, src, dest string) (bool, error) {
	return func(srcinfo os.FileInfo, src, dest string) (bool, error) {
		//slog.Debug("Checking file", "file", src)
		if srcinfo.IsDir() {
			//slog.Debug("Skipping directory", "dir", src)
			return false, nil
		}
		if filepath.Ext(src) != ext {
			slog.Debug("Skipping file", "file", src, "ext", filepath.Ext(src))
			return true, nil
		}
		return false, nil
	}
}
