package console

import (
	"github.com/renan5g/goframework/console/console"
	consolecontract "github.com/renan5g/goframework/contracts/console"
	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goravel.console"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewConsole(), nil
	})
}

func (receiver *ServiceProvider) Boot(app core.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app core.Application) {
	artisan := app.MakeArtisan()
	config := app.MakeConfig()
	artisan.Register([]consolecontract.Command{
		console.NewListCommand(artisan),
		console.NewKeyGenerateCommand(config),
		console.NewMakeCommand(),
	})
}
