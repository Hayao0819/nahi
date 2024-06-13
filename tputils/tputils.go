package tputils

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func ApplyTemplate(file string, data any) (*bytes.Buffer, error) {
	return ApplyToFile(file, data)
}

func ApplyToFile(file string, data any) (*bytes.Buffer, error) {
	tmpl, err := template.New(filepath.Base(file)).ParseFiles(file)
	if err != nil {
		return nil, err
	}
	return Apply(tmpl, data)
}

func ApplyToText(text string, data any) (*bytes.Buffer, error) {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		return nil, err
	}
	return Apply(tmpl, data)
}

func Apply(tmpl *template.Template, data any) (*bytes.Buffer, error) {
	writer := bytes.NewBufferString("")
	if err := tmpl.Execute(writer, data); err != nil {
		return nil, err
	}
	return writer, nil
}
