package futils

import "os"

// 文字列が正常なディレクトリへのパスかどうかを確認します
func IsDir(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

// 文字列が正常なファイルパスかどうかを調べます
func IsFile(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}

// シンボリックリンクかどうか
// 参考: https://github.com/eihigh/filetest
// Thanks eihigh <eihigh.contact@gmail.com>
func IsSymlink(path string) bool {
	stat, err := os.Lstat(path)
	return err == nil && stat.Mode()&os.ModeSymlink == os.ModeSymlink
}

// ファイルが存在するかどうか
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
