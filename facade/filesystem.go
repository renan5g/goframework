package facade

import "github.com/renan5g/goframework/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
