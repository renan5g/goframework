package fiber

import (
	"github.com/renan5g/goframework/contracts/config"
	"github.com/renan5g/goframework/contracts/core"
	"github.com/renan5g/goframework/contracts/http"
	"github.com/renan5g/goframework/contracts/log"
	"github.com/renan5g/goframework/contracts/validation"
)

const RouteBinding = "goframework.fiber.route"

var (
	App              core.Application
	ConfigFacade     config.Config
	LogFacade        log.Log
	ValidationFacade validation.Validation
	ViewFacade       http.View
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app core.Application) {
	App = app

	app.BindWith(RouteBinding, func(app core.Application, parameters map[string]any) (any, error) {
		return NewRoute(app.MakeConfig(), parameters)
	})
}

func (receiver *ServiceProvider) Boot(app core.Application) {
	ConfigFacade = app.MakeConfig()
	LogFacade = app.MakeLog()
	ValidationFacade = app.MakeValidation()
	ViewFacade = app.MakeView()

	app.Publishes("github.com/renan5g/goframework/pkg/fiber", map[string]string{
		"config/cors.go": app.ConfigPath("cors.go"),
	})
}
