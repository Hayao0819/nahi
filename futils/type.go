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
