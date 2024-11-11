package nconf

import (
	"os"
	"path/filepath"

	"github.com/go-viper/mapstructure/v2"
)

func (n *Nconf[T]) AddLoadDir(dirs ...string) {
	n.dirs = append(n.dirs, dirs...)
}

func (n *Nconf[T]) AddLoadFile(filenames ...string) {
	n.filenames = append(n.filenames, filenames...)
}

func (n *Nconf[T]) files() []string {
	files := []string{}
	for _, dir := range n.dirs {
		for _, filename := range n.filenames {
			files = append(files, filepath.Join(dir, filename))
		}
	}
	return files
}

func (n *Nconf[T]) Load() error {
	merged := map[string]string{}
	for _, file := range n.files() {
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		d, err := parseKeyValue(f)
		if err != nil {
			return err
		}

		for k, v := range d {
			merged[k] = v
		}

		f.Close()
	}

	// map[string]stringをいい感じにTに変換する
	if err := mapstructure.Decode(merged, &n.data); err != nil {
		return err
	}

	n.loaded = true
	return nil
}
