package log

import (
	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goframework.log"

type ServiceProvider struct {
}

func (log *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (log *ServiceProvider) Boot(app core.Application) {

}
