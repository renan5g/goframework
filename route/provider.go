package route

import (
	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goravel.route"

type ServiceProvider struct {
}

func (route *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewRoute(app.MakeConfig()), nil
	})
}

func (route *ServiceProvider) Boot(app core.Application) {

}
