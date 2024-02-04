package filesystem

import (
	"path/filepath"
	"strings"

	"github.com/renan5g/goframework/contracts/filesystem"
	"github.com/renan5g/goframework/support/file"
)

func fullPathOfFile(filePath string, source filesystem.File, name string) (string, error) {
	extension := filepath.Ext(name)
	if extension == "" {
		var err error
		extension, err = file.Extension(source.File(), true)
		if err != nil {
			return "", err
		}

		return filepath.Join(filePath, strings.TrimSuffix(strings.TrimPrefix(filepath.Base(name), string(filepath.Separator)), string(filepath.Separator))+"."+extension), nil
	} else {
		return filepath.Join(filePath, strings.TrimPrefix(filepath.Base(name), string(filepath.Separator))), nil
	}
}

func validPath(path string) string {
	realPath := strings.TrimPrefix(path, "./")
	realPath = strings.TrimPrefix(realPath, "/")
	realPath = strings.TrimPrefix(realPath, ".")
	if realPath != "" && !strings.HasSuffix(realPath, "/") {
		realPath += "/"
	}

	return realPath
}
