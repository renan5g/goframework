package redis

import (
	"context"

	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goframework.redis"

var App core.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app core.Application) {
	App = app

	app.BindWith(Binding, func(app core.Application, parameters map[string]any) (any, error) {
		return NewRedis(context.Background(), app.MakeConfig(), parameters["store"].(string))
	})
}

func (receiver *ServiceProvider) Boot(app core.Application) {

}
