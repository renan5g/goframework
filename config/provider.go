package config

import (
	"github.com/renan5g/goframework/contracts/core"
	"github.com/renan5g/goframework/support"
)

const Binding = "goframework.config"

type ServiceProvider struct {
}

func (config *ServiceProvider) Register(app core.Application) {
	app.Singleton(Binding, func(app core.Application) (any, error) {
		return NewConfig(support.EnvPath), nil
	})
}

func (config *ServiceProvider) Boot(app core.Application) {

}
