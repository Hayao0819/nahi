package tputils

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func ApplyTemplate(file string, data any) (*bytes.Buffer, error) {
	tmpl, err := template.New(filepath.Base(file)).ParseFiles(file)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBufferString("")
	if err := tmpl.Execute(writer, data); err != nil {
		return nil, err
	}
	return writer, nil
}
