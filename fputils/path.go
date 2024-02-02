package fputils

import (
	"os"
	"path/filepath"
	"strings"
)

// 拡張子を除いたファイル名を返します
func BaseWithoutExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// ~/を置き換え
func ReplaceTilde(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", home, 1), nil
}
