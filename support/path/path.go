package path

import (
	"github.com/renan5g/goframework/facade"
)

func App(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().Path(finalPath)
}

func Base(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().BasePath(finalPath)
}

func Config(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().ConfigPath(finalPath)
}

func Database(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().DatabasePath(finalPath)
}

func Storage(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().StoragePath(finalPath)
}

func Public(paths ...string) string {
	finalPath := ""
	if len(paths) >= 1 {
		finalPath = paths[0]
	}

	return facade.App().PublicPath(finalPath)
}
