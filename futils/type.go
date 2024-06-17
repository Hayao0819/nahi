package futils

import (
	"net/http"
	"os"
)

func DetectFileType(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "application/octet-stream", err
	}
	return http.DetectContentType(bytes), nil
}

func Executable(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return false, err
	}

	return fi.Mode().IsRegular() && (fi.Mode().Perm()&0111) != 0, nil
}
