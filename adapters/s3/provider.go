package s3

import (
	"context"

	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goframework.s3"

var App core.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app core.Application) {
	App = app

	app.BindWith(Binding, func(app core.Application, parameters map[string]any) (any, error) {
		return NewS3(context.Background(), app.MakeConfig(), parameters["disk"].(string))
	})
}

func (receiver *ServiceProvider) Boot(app core.Application) {
	app.Publishes("github.com/renan5g/goframework/pkg/s3", map[string]string{
		"config/s3.go": app.ConfigPath(""),
	})
}
