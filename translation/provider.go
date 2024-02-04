package translation

import (
	"context"
	"path/filepath"

	"github.com/renan5g/goframework/contracts/core"
)

const Binding = "goravel.translation"

type ServiceProvider struct {
}

func (translation *ServiceProvider) Register(app core.Application) {
	app.BindWith(Binding, func(app core.Application, parameters map[string]any) (any, error) {
		config := app.MakeConfig()
		logger := app.MakeLog()
		locale := config.GetString("app.locale")
		fallback := config.GetString("app.fallback_locale")
		loader := NewFileLoader([]string{filepath.Join("lang")})
		trans := NewTranslator(parameters["ctx"].(context.Context), loader, locale, fallback, logger)
		trans.SetLocale(locale)
		return trans, nil
	})
}

func (translation *ServiceProvider) Boot(app core.Application) {

}
