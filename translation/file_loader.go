package translation

import (
	"os"
	"path/filepath"

	"github.com/renan5g/goframework/support/file"
	"github.com/renan5g/goframework/support/json"
)

type FileLoader struct {
	paths []string
}

func NewFileLoader(paths []string) *FileLoader {
	return &FileLoader{
		paths: paths,
	}
}

func (f *FileLoader) Load(locale string, group string) (map[string]any, error) {
	for _, path := range f.paths {
		var val map[string]any
		fullPath := filepath.Join(path, locale, group+".json")
		if group == "*" {
			fullPath = filepath.Join(path, locale+".json")
		}

		if file.Exists(fullPath) {
			data, err := os.ReadFile(fullPath)
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(data, &val); err != nil {
				return nil, err
			}
			return val, nil
		}
	}
	return nil, ErrFileNotExist
}
