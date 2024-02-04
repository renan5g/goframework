package filesystem

import (
	configcontract "github.com/renan5g/goframework/contracts/config"
	"github.com/renan5g/goframework/contracts/core"
	filesystemcontract "github.com/renan5g/goframework/contracts/filesystem"
)

const Binding = "goframework.filesystem"

var ConfigFacade configcontract.Config
var StorageFacade filesystemcontract.Storage

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewStorage(app.MakeConfig()), nil
	})
}

func (database *ServiceProvider) Boot(app core.Application) {
	ConfigFacade = app.MakeConfig()
	StorageFacade = app.MakeStorage()
}
