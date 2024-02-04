package http

import (
	"github.com/renan5g/goframework/contracts/cache"
	"github.com/renan5g/goframework/contracts/config"
	consolecontract "github.com/renan5g/goframework/contracts/console"
	"github.com/renan5g/goframework/contracts/core"
	"github.com/renan5g/goframework/contracts/http"
	"github.com/renan5g/goframework/http/console"
)

const BindingRateLimiter = "goravel.rate_limiter"
const BindingView = "goravel.view"

type ServiceProvider struct{}

var (
	CacheFacade       cache.Cache
	ConfigFacade      config.Config
	RateLimiterFacade http.RateLimiter
)

func (http *ServiceProvider) Register(app core.Application) {
	app.Singleton(BindingRateLimiter, func(app core.Application) (any, error) {
		return NewRateLimiter(), nil
	})
	app.Singleton(BindingView, func(app core.Application) (any, error) {
		return NewView(), nil
	})
}

func (http *ServiceProvider) Boot(app core.Application) {
	CacheFacade = app.MakeCache()
	ConfigFacade = app.MakeConfig()
	RateLimiterFacade = app.MakeRateLimiter()

	http.registerCommands(app)
}

func (http *ServiceProvider) registerCommands(app core.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
