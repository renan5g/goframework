package validation

import (
	consolecontract "github.com/renan5g/goframework/contracts/console"
	"github.com/renan5g/goframework/contracts/core"
	"github.com/renan5g/goframework/validation/console"
)

const Binding = "goravel.validation"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewValidation(), nil
	})
}

func (database *ServiceProvider) Boot(app core.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app core.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
