package cache

import (
	"github.com/renan5g/goframework/cache/console"
	contractsconsole "github.com/renan5g/goframework/contracts/console"
	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goframework.cache"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		config := app.MakeConfig()
		log := app.MakeLog()
		store := config.GetString("cache.default")

		return NewCacher(config, log, store)
	})
}

func (database *ServiceProvider) Boot(app core.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app core.Application) {
	app.MakeArtisan().Register([]contractsconsole.Command{
		console.NewClearCommand(app.MakeCache()),
	})
}
