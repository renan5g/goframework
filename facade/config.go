package facade

import (
	"github.com/renan5g/goframework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
